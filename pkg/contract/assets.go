package contract

import (
	"context"
	"wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
)

type operations struct {
	ops []*grpc.ContractAssetOperation
}

type Operations interface {
	Append(operation *grpc.ContractAssetOperation)

	GetResult() []*grpc.ContractAssetOperation
}

func NewOperations() Operations {
	return &operations{
		ops: make([]*grpc.ContractAssetOperation, 0),
	}
}

func (o *operations) Append(operation *grpc.ContractAssetOperation) {
	o.ops = append(o.ops, operation)
}

func (o *operations) GetResult() []*grpc.ContractAssetOperation {
	return o.ops
}

type ExecutionContext struct {
	Ctx context.Context

	Client *contract2.ContractServiceClient

	State      State
	Operations Operations

	Tx *contract2.ContractTransaction

	nonce int32
}

func NewExecutionContext(ctx context.Context, cl *contract2.ContractServiceClient, state2 State, ops Operations, tx *contract2.ContractTransaction) *ExecutionContext {
	return &ExecutionContext{
		Ctx:        ctx,
		Client:     cl,
		State:      state2,
		Operations: ops,
		Tx:         tx,
		nonce:      0,
	}
}

func (c ExecutionContext) GetNonce() int32 {
	return c.nonce + 1
}
