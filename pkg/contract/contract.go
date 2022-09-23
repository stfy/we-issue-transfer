package contract

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	grpc "wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
)

type contract struct {
	c contract2.ContractServiceClient

	authToken string
	handlers  map[string]*Action
}

type Contract interface {
	HandleTx(tx *contract2.ContractTransaction)

	Auth(t string)

	Handler(t Action)
}

const actionkey = "action"

func NewContract(c contract2.ContractServiceClient) Contract {
	return &contract{
		c:         c,
		authToken: "",
		handlers:  map[string]*Action{},
	}
}

func (c *contract) Handler(t Action) {
	c.handlers[t.Name()] = &t
}

func (c *contract) Auth(t string) {
	c.authToken = t
}

func (c *contract) Ctx() context.Context {
	md := metadata.Pairs("authorization", c.authToken)

	return metadata.NewOutgoingContext(context.Background(), md)
}

func (c *contract) HandleTx(tx *contract2.ContractTransaction) {
	switch tx.Type {
	case 103:
		resp, err := c.c.CommitExecutionSuccess(c.Ctx(), &contract2.ExecutionSuccessRequest{
			TxId:            tx.Id,
			Results:         nil,
			AssetOperations: nil,
		})

		if err != nil {
			panic(err)
		}

		zap.S().Infoln("Contract successfully created", resp)
	case 104:
		action, err := txAction(tx)

		zap.S().Infoln("Try to execute action", action)
		zap.S().Infoln("Available action handlers", c.handlers)

		ec := NewExecutionContext(c.Ctx(), &c.c, NewState(c.c, tx.ContractId), NewOperations(), tx)
		handler := *c.handlers[action]

		if handler == nil {
			zap.S().Infoln("Action not founded")

			_, _ = c.c.CommitExecutionError(c.Ctx(), &contract2.ExecutionErrorRequest{
				TxId:    tx.Id,
				Message: "Action not founded",
				Code:    1,
			})
			return
		}

		handler.Handle(*ec, tx)

		if err != nil {
			c.c.CommitExecutionError(c.Ctx(), &contract2.ExecutionErrorRequest{
				TxId:    tx.Id,
				Message: err.Error(),
				Code:    0,
			})

			break
		}

		zap.S().Infoln("Action executed", tx.Id, tx.Params)
		zap.S().Infoln("Entries changes", ec.State.GetEntries())
		zap.S().Infoln("Asset operations", ec.Operations.GetResult())

		results := ec.State.GetEntries()
		ops := ec.Operations.GetResult()

		var header, trailer metadata.MD

		resp, err := c.c.CommitExecutionSuccess(
			c.Ctx(),
			&contract2.ExecutionSuccessRequest{
				TxId:            tx.Id,
				Results:         results,
				AssetOperations: ops,
			},
			g.Header(&header),
			g.Trailer(&trailer),
		)

		if err != nil {
			zap.S().Errorln("header", header)
			zap.S().Errorln("trailer", trailer)
			zap.S().Errorln("err", err)

			_, _ = c.c.CommitExecutionError(c.Ctx(), &contract2.ExecutionErrorRequest{
				TxId:    tx.Id,
				Message: err.Error(),
				Code:    0,
			})
		} else {
			zap.S().Infoln("Contract action successfully executed", resp, err)
		}
	}
}

func txAction(tx *contract2.ContractTransaction) (string, error) {
	index := slices.IndexFunc(tx.Params, func(tx *grpc.DataEntry) bool {
		return tx.Key == actionkey
	})

	if index == -1 {
		return "", errors.New("param \"action\" not found ")
	}

	p := tx.Params[index]

	return p.Value.(*grpc.DataEntry_StringValue).StringValue, nil
}

func TxParam(tx *contract2.ContractTransaction, key string) (*grpc.DataEntry, error) {
	index := slices.IndexFunc(tx.Params, func(tx *grpc.DataEntry) bool {
		return tx.Key == key
	})

	if index == -1 {
		return nil, errors.New("param not found")
	}

	return tx.Params[index], nil
}
