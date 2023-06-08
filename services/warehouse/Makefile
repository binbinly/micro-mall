GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	@go install github.com/envoyproxy/protoc-gen-validate

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/health.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. --validate_out="lang=go:." proto/warehouse.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o warehouse *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t warehouse:latest .

.PHONY: doc
doc:
	protoc --doc_out=. --doc_opt=markdown,README.md proto/*.proto