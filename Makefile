SHELL = /bin/bash
MAKEFLAGS += --silent

.PHONY: build
build: api-clean api
	go build -o ./build/users-api-server ./cmd/users-api-server/main.go

build-clean:
	rm -rf ./build/*

run-server:
	go run cmd/users-api-server/main.go

.PHONY: api
api:
	oapi-codegen --config=./api/openapi-specs/configs/server.yaml \
		api/openapi-specs/openapi.yaml > internal/platform/server/api_server.gen.go

api-clean:
	find internal/platform/server -type f  -name '*.gen.go' -delete

lint:
	gofmt -w -s .
	goimports -w .
	golangci-lint run

clean: build-clean api-clean

all: clean api build

test:
	go test ./...

test-coverage:
	go test ./... -cover

test-coverage-reporter:
	scripts/test/reporter.sh
