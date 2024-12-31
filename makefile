proto:
	 protoc \
	   --go_out=./proto/pb \
	   --go_opt=paths=source_relative \
	   --go-grpc_out=./proto/pb \
	   --go-grpc_opt=paths=source_relative \
	   --proto_path=./proto \
	   ./proto/**/*.proto


.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go run cmd/server/main.go
.PHONY:build