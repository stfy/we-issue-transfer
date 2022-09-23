package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"os"
	"strings"
	"wego/internal/common"
	"wego/pkg/contract"
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
	common.SetupLogger("DEV")

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
		_ = c.State.SetInt("my-test_key_int", 200)
	}))

	zap.S().Infoln("Contract started")

	for {
		msg, err := ccl.Recv()

		zap.S().Infoln("Received rpc message", msg)

		if err != nil {
			panic(err)
		}

		c.Auth(msg.AuthToken)
		c.HandleTx(msg.Transaction)
	}
}
