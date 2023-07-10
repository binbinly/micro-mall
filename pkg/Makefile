GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	go install github.com/golang/protobuf/proto@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-micro/cli/cmd/go-micro@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: proto
proto:
	protoc --proto_path=./proto \
    --micro_out ./proto --micro_opt paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./proto" \
    --go_out ./proto --go_opt paths=source_relative \
    proto/core/\*.proto