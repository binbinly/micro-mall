# Go parameters
PROJECT_NAME := "gateway"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /example)

GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: dev
dev:
	go run main.go server -c local.yaml

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/member.proto

.PHONY: proto
proto:
	protoc -I ./proto \
	  --go_out=paths=source_relative:./proto \
	  --validate_out="lang=go,paths=source_relative:./proto" \
	  --go-grpc_out=paths=source_relative:./proto \
	  --grpc-gateway_out=paths=source_relative:./proto \
	  --openapiv2_out ./docs \
      --openapiv2_opt logtostderr=true \
      --openapiv2_opt use_go_templates=true \
	  ./proto/market/*.proto

.PHONY: build
build:
	go build -o member *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t member:latest

#检查代码规范
.PHONY: lint
lint:
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}


.PHONY: common
common:
	protoc -I ./proto --go_out=paths=source_relative:./proto  --validate_out="lang=go,paths=source_relative:./proto" ./proto/common/*.proto