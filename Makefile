SHELL = /bin/bash

api-clean:
	scripts/api/clean.sh

api-codegen: api-clean
	scripts/api/codegen.sh

build-generate: api-codegen
	go build -o ./build/api ./cmd/api/main.go

build-clean:
	rm -f ./build/*

lint:
	golangci-lint run
