build_proto:
	protoc --go_out=. --go-grpc_out=. cofee_shop.proto