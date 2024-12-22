run:
	go run cmd/main.go
proto:
	protoc --go_out=./internal/common --go_opt=paths=source_relative \
		--go-grpc_out=./internal/common --go-grpc_opt=paths=source_relative \
		shop.proto
