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

api-lint:
	vacuum lint -d -n error -r api/openapi-specs/configs/ruleset.yaml api/openapi-specs/openapi.yaml

lint:
	golangci-lint run

clean: build-clean api-clean

all: clean api build

tests:
	go test ./...

tests-coverage:
	go test ./... -cover

tests-coverage-reporter:
	scripts/test/reporter.sh

install-tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/vektra/mockery/v2@v2.33.2
	go install github.com/onsi/ginkgo/v2/ginkgo@2.12.0
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.14.0
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
