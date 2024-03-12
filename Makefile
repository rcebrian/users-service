SHELL = /bin/bash
MAKEFLAGS += --silent

COVERAGE_DIR=coverage
GOLANGCI_LINT_VERSION=v1.56.2

all: clean api build

run:
	godotenv -f .env go run cmd/main.go

.PHONY: build
build:
	go build -o ./build/users-api-server ./cmd/users-api-server/main.go

.PHONY: api
api:
	oapi-codegen --config=./api/openapi-specs/configs/server.yaml \
		api/openapi-specs/openapi.yaml > internal/platform/server/api_server.gen.go

lint:
	golangci-lint run
	vacuum lint -d -n error -r api/openapi-specs/configs/ruleset.yaml api/openapi-specs/openapi.yaml

clean:
	rm -rf ./build/*

test:
	go test ./... -coverprofile=${COVERAGE_DIR}/coverage.out

.PHONY: coverage
coverage: test
	go tool cover -html=${COVERAGE_DIR}/coverage.out -o ${COVERAGE_DIR}/coverage.html
	go tool cover -func=${COVERAGE_DIR}/coverage.out > ${COVERAGE_DIR}/coverage.txt

install-tools:
	# run
	go install github.com/joho/godotenv/cmd/godotenv@latest
	# code generation
	go install github.com/vektra/mockery/v2@v2.42.0
	go install github.com/daveshanley/vacuum@v0.9.10
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v2.1.0
	# lint tools
	go install golang.org/x/tools/cmd/goimports@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin ${GOLANGCI_LINT_VERSION}
