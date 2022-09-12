package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"os"
	"strings"
	"wego/pkg/contract"
	"wego/pkg/contract/we"
	wecontract "wego/pkg/grpc/contract"
)

func buildUri(node string, nodePort string) string {
	var uri strings.Builder
	uri.WriteString(node)
	uri.WriteString(":")
	uri.WriteString(nodePort)

	return uri.String()
}

func main() {
	var (
		connectionId    = os.Getenv("CONNECTION_ID")
		connectionToken = os.Getenv("CONNECTION_TOKEN")
		node            = os.Getenv("NODE")
		nodePort        = os.Getenv("NODE_PORT")
	)

	uri := buildUri(node, nodePort)

	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	md := metadata.Pairs("authorization", connectionToken)

	if err != nil {
		panic(err)
	}

	client := wecontract.NewContractServiceClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ccl, err := client.Connect(ctx, &wecontract.ConnectionRequest{
		ConnectionId: connectionId,
	})

	if err != nil {
		panic(err)
	}

	c := contract.NewContract(client)

	c.Handler(contract.NewContractAction("set-key", func(c contract.ExecutionContext, tx *wecontract.ContractTransaction) {
		_ = c.State.SetString("my-test_key", "value of my string key")
		_ = c.State.SetInt("my-test_key_int", 228)
	}))

	c.Handler(contract.NewContractAction("issue", func(c contract.ExecutionContext, tx *wecontract.ContractTransaction) {
		id, nonce, _ := we.CalculateAssetId(c)
		t := we.NewAssetFromId(id)

		_ = t.Issue(c, "weUSD", "USDT token", 8, 100_000, false, nonce)

		var k strings.Builder
		k.WriteString("issued")
		k.WriteString(":")
		k.WriteString(id)

		_ = c.State.SetInt(k.String(), int64(100_000))
	}))

	c.Handler(contract.NewContractAction("transfer", func(c contract.ExecutionContext, tx *wecontract.ContractTransaction) {
		a, _ := contract.TxParam(tx, "asset")
		r, _ := contract.TxParam(tx, "recipient")

		am := 50_000

		token := we.NewAssetFromId(a.GetStringValue())
		_ = token.Transfer(c, r.GetStringValue(), int64(am))

		var k strings.Builder
		k.WriteString("transfered")
		k.WriteString(":")
		k.WriteString(a.GetStringValue())
		k.WriteString(":")
		k.WriteString(r.String())

		_ = c.State.SetInt(k.String(), int64(am))
	}))

	for {
		msg, err := ccl.Recv()

		if err != nil {
			panic(err)
		}

		c.Auth(msg.AuthToken)
		c.HandleTx(msg.Transaction)
	}
}
