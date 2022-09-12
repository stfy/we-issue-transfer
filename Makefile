build:
	docker build -f ./Dockerfile . -t issue-transfer:0.0.1
pb:
	@protoc --proto_path=protobuf  \
			--go_out=pkg/grpc --go_opt=paths=source_relative \
			--go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative \
			protobuf/data_entry.proto \
			protobuf/*.proto \
			protobuf/address/*.proto \
			protobuf/alias/*.proto \
			protobuf/entity/*.proto \
			protobuf/contract/*.proto \
			protobuf/contract_asset_operation/*.proto \
			protobuf/genesis/*.proto \
			protobuf/managed/*.proto \
			protobuf/messagebroker/*.proto \
			protobuf/pki/*.proto \
			protobuf/privacy/*.proto \
			protobuf/transaction/*.proto \
			protobuf/util/*.proto

