proto:
	protoc ./proto/*.proto --go_out=./ --go-grpc_out=./

.PHONY: proto