GOPATH:=$(shell go env GOPATH)

.PHONY: dev
dev:
	@go run main.go --registry consul --registry_address 127.0.0.1:8500

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t order:latest .