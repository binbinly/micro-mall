GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --proto_path=C:\\Users\\HAILAN\\AppData\\Local\\JetBrains\\GoLand2022.3\\protoeditor --micro_out=. --go_out=:. --validate_out="lang=go:." proto/center.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o center *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t center:latest .