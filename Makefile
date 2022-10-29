SHELL = /bin/bash
MAKEFLAGS += --silent

.PHONY: build
build:
	go build -o ./build/api ./cmd/api/main.go

build-clean:
	rm -f ./build/*

.PHONY: api
api:
	scripts/api/codegen.sh

api-clean:
	find internal/platform/server/openapi -type f -not -name '*_service.go' -delete

lint:
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
