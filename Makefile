.PHONY: build
build:
	go build -v ./cmd/kiddy-line-processor

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: generate
generate:
	./protoc/bin/protoc --proto_path=proto --go_out=plugins=grpc:proto lineProcessor.proto

.PHONY: lint
lint:
	#docker-compose up

.PHONY: run
run:
	docker-compose up

.PHONY: stop
stop:
	docker-compose down

.DEFAULT_GOAL := build

