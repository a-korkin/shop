run_api:
	go run cmd/main.go -a
run_grpc:
	go run cmd/main.go -g
proto:
	protoc --go_out=./internal/common --go_opt=paths=source_relative \
		--go-grpc_out=./internal/common --go-grpc_opt=paths=source_relative \
		shop.proto
