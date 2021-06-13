gen:
	protoc  --proto_path=api api/*.proto \
		--go_out=. --go-grpc_out=require_unimplemented_servers=false:. \
		--grpc-gateway_out=:.

run-grpc:
	go run cmd/*.go

run-http:
	go run cmd/http/*.go