
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	mkdir docs

.PHONY: proto
proto:
	protoc --proto_path=./proto \
    --micro_out ./proto --micro_opt paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./proto" \
    --go_out ./proto --go_opt paths=source_relative \
    proto/core/\*.proto