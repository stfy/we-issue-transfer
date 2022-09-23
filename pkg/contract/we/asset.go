package we

import (
	"go.uber.org/zap"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"wego/pkg/contract"
	"wego/pkg/grpc"
	contract2 "wego/pkg/grpc/contract"
	"wego/pkg/grpc/contract_asset_operation"
)

type asset struct {
	id string
}

type Asset interface {
	Transfer(ec contract.ExecutionContext, recipient string, amount int64) error

	Issue(ec contract.ExecutionContext, name string, description string, decimals int32, qty int64, isreissuable bool, nonce int32) error
}

func CalculateAssetId(ec contract.ExecutionContext) (string, int32, error) {
	c := *ec.Client

	nonce := ec.GetNonce()

	zap.S().Infoln("CalculateAssetId", ec, "nonce", nonce)

	var header, trailer metadata.MD

	id, err := c.CalculateAssetId(ec.Ctx,
		&contract2.CalculateAssetIdRequest{Nonce: nonce},
		g.Header(&header),
		g.Trailer(&trailer),
	)

	if err != nil {
		zap.S().Errorln("CalculateAssetId", "header", header)
		zap.S().Errorln("CalculateAssetId", "trailer", trailer)
		zap.S().Errorln("CalculateAssetId", "err", err)

		return "", 0, err
	}

	return id.Value, nonce, nil
}

func (a *asset) Transfer(ec contract.ExecutionContext, recipient string, amount int64) error {
	zap.S().Infoln("Transfer", "context", ec)
	zap.S().Infoln("Transfer",
		"asset", a.id,
		"recipient", recipient,
		"amount", amount,
	)

	op := &contract_asset_operation.ContractTransferOut{
		AssetId:   wrapperspb.String(a.id),
		Recipient: recipient,
		Amount:    amount,
	}

	assetOperation := grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractTransferOut{ContractTransferOut: op}}

	ec.Operations.Append(&assetOperation)

	return nil
}

func (a *asset) Issue(ec contract.ExecutionContext, name string, description string, decimals int32, qty int64, isreissuable bool, nonce int32) error {
	op := &contract_asset_operation.ContractIssue{
		AssetId:      a.id,
		Name:         name,
		Description:  description,
		Decimals:     decimals,
		IsReissuable: isreissuable,
		Quantity:     qty,
		Nonce:        nonce,
	}

	assetOperation := &grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractIssue{ContractIssue: op}}

	ec.Operations.Append(assetOperation)

	return nil
}

func (a *asset) Reissue(ec contract.ExecutionContext, qty int64, isreissuable bool) error {
	zap.S().Infoln("Reissue", "context", ec)
	zap.S().Infoln("Reissue",
		"asset", a.id,
		"qty", qty,
		"isreissuable", isreissuable,
	)

	op := &contract_asset_operation.ContractReissue{
		AssetId:      a.id,
		Quantity:     qty,
		IsReissuable: isreissuable,
	}

	assetOperation := grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractReissue{ContractReissue: op}}

	ec.Operations.Append(&assetOperation)

	return nil
}

func (a *asset) Burn(ec contract.ExecutionContext, qty int64) error {
	zap.S().Infoln("Burn", "context", ec)
	zap.S().Infoln("Burn",
		"asset", a.id,
		"qty", qty,
	)

	op := &contract_asset_operation.ContractBurn{
		AssetId: wrapperspb.String(a.id),
		Amount:  qty,
	}

	assetOperation := grpc.ContractAssetOperation{Operation: &grpc.ContractAssetOperation_ContractBurn{ContractBurn: op}}

	ec.Operations.Append(&assetOperation)

	return nil
}
