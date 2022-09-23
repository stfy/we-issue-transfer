package contract

import (
	"context"
	"go.uber.org/zap"
	"wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
)

type State interface {
	SetInt(key string, value int64) error
	SetString(key string, value string) error

	Get(ctx context.Context, key string) (*grpc.DataEntry, error)
	GetEntries() []*grpc.DataEntry
}

type state struct {
	c contract2.ContractServiceClient

	contractId string

	entries []*grpc.DataEntry
}

func NewState(c contract2.ContractServiceClient, contract string) State {
	return &state{c: c,
		contractId: contract,
		entries:    make([]*grpc.DataEntry, 0),
	}
}

func (r *state) GetEntries() []*grpc.DataEntry {
	return r.entries
}

func (r *state) SetString(key string, v string) error {
	r.entries = append(r.entries, &grpc.DataEntry{Key: key, Value: &grpc.DataEntry_StringValue{StringValue: v}})

	return nil
}

func (r *state) SetInt(key string, v int64) error {
	r.entries = append(r.entries, &grpc.DataEntry{Key: key, Value: &grpc.DataEntry_IntValue{IntValue: v}})
	return nil
}

func (r *state) Get(ctx context.Context, key string) (*grpc.DataEntry, error) {
	resp, err := r.c.GetContractKey(ctx, &contract2.ContractKeyRequest{
		ContractId: r.contractId,
		Key:        key,
	})

	return resp.Entry, err
}

type contractAction struct {
	name string
	fn   func(ExecutionContext, *contract2.ContractTransaction)
}

type Action interface {
	Name() string
	Handle(ec ExecutionContext, tx *contract2.ContractTransaction)
}

func NewContractAction(name string, handler func(ExecutionContext, *contract2.ContractTransaction)) Action {
	return &contractAction{name: name, fn: handler}
}

func (r *contractAction) Name() string {
	return r.name
}

func (r *contractAction) Handle(ec ExecutionContext, tx *contract2.ContractTransaction) {
	r.fn(ec, tx)

	zap.S().Infoln("Executed handler", r.name, ec, tx)
}
