SHELL = /bin/bash
MAKEFLAGS += --silent

.PHONY: build
build: api
	go build -o ./build/api ./cmd/api/main.go

build-clean:
	rm -f ./build/*

.PHONY: api
api: api-clean
	scripts/api/codegen.sh

api-clean:
	find internal/platform/server/openapi -type f -not -name '*_service.go' -delete

lint:
	golangci-lint run
