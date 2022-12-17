proto:
	protoc ./proto/calculator.proto --go_out=./ --go-grpc_out=./

.PHONY: proto