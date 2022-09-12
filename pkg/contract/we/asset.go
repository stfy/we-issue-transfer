package we

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"wego/pkg/contract"
	"wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
	"wego/pkg/grpc/contract_asset_operation"
)

type asset struct {
	id []byte
}

type Asset interface {
	Transfer(ec contract.ExecutionContext, recipient string, amount int64) error

	Issue(ec contract.ExecutionContext, name string, description string, decimals int32, qty int64, isreissuable bool, nonce int32) error
}

func CalculateAssetId(ec contract.ExecutionContext) (string, int32, error) {
	c := *ec.Client

	nonce := ec.GetNonce()

	id, err := c.CalculateAssetId(ec.Ctx, &contract2.CalculateAssetIdRequest{Nonce: nonce})

	if err != nil {
		return "", 0, err
	}

	return id.Value, nonce, nil
}

func NewAssetFromId(id string) Asset {
	return &asset{id: []byte(id)}
}

func (a *asset) Transfer(ec contract.ExecutionContext, recipient string, amount int64) error {
	op := &contract_asset_operation.ContractTransferOut{
		AssetId:   wrapperspb.Bytes(a.id),
		Recipient: []byte(recipient),
		Amount:    amount,
	}

	assetOperation := grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractTransferOut{ContractTransferOut: op}}

	ec.Operations.Append(&assetOperation)

	return nil
}

func (a *asset) Issue(ec contract.ExecutionContext, name string, description string, decimals int32, qty int64, isreissuable bool, nonce int32) error {
	op := &contract_asset_operation.ContractIssue{
		AssetId:      a.id,
		Name:         []byte(name),
		Description:  []byte(description),
		Decimals:     decimals,
		IsReissuable: isreissuable,
		Quantity:     qty,
		Nonce:        nonce,
	}

	assetOperation := &grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractIssue{ContractIssue: op}}

	ec.Operations.Append(assetOperation)

	return nil
}
