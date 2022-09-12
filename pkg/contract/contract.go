package contract

import (
	"context"
	"errors"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/metadata"
	"log"
	"wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
)

type contract struct {
	c contract2.ContractServiceClient

	authToken string
	handlers  map[string]Action
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
		handlers:  map[string]Action{},
	}
}

func (c *contract) Handler(t Action) {
	c.handlers[t.Name()] = t
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

		log.Println("Contract successfully created", resp)
	case 104:
		action, err := txAction(tx)

		if err != nil {

			c.c.CommitExecutionError(c.Ctx(), &contract2.ExecutionErrorRequest{
				TxId:    tx.Id,
				Message: err.Error(),
				Code:    0,
			})

			break
		}

		state := NewState(c.c, tx.ContractId)

		ec := NewExecutionContext(c.Ctx(), &c.c, state, NewOperations(), tx)

		c.handlers[action].Handle(*ec, tx)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		log.Println("Action executed", tx.Id, tx.Params)
		log.Println(ec.State.GetEntries())
		log.Println(ec.Operations.GetResult())

		results := ec.State.GetEntries()
		ops := ec.Operations.GetResult()

		resp, err := c.c.CommitExecutionSuccess(c.Ctx(), &contract2.ExecutionSuccessRequest{
			TxId:            tx.Id,
			Results:         results,
			AssetOperations: ops,
		})

		if err != nil {
			_, _ = c.c.CommitExecutionError(c.Ctx(), &contract2.ExecutionErrorRequest{
				TxId:    tx.Id,
				Message: err.Error(),
				Code:    0,
			})
		} else {
			log.Println("Contract action successfully executed", resp, err)
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
