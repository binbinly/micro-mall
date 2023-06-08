
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi
	go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	mkdir docs

.PHONY: api
api:
	protoc --openapiv2_out=./doc --proto_path=. proto/*/*.proto

.PHONY: doc
doc:
	protoc -I ./proto --doc_out=. --doc_opt=markdown,README.md proto/*/*.proto

.PHONY: proto
proto:
	protoc --proto_path=./proto --proto_path=C:\\Users\\HAILAN\\AppData\\Local\\JetBrains\\GoLand2022.3\\protoeditor \
    --micro_out ./proto --micro_opt paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./proto" \
    --go_out ./proto --go_opt paths=source_relative \
    proto/core/\*.proto