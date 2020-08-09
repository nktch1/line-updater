.PHONY: build
build:
	go build -v ./cmd/lineProcessor -o lineProcessor

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: generate
generate:
	protoc1 --proto_path=proto --go_out=plugins=grpc:./pkg/rpcserver lineProcessor.proto

.PHONY: lint
lint:
	golint ./...

.PHONY: run
run:
	docker-compose up

.PHONY: stop
stop:
	docker-compose down

.DEFAULT_GOAL := build