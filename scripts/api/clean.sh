#!/usr/bin/env bash
: <<DOCS
Remove the generated code by OpenAPI Generator
DOCS

find internal/platform/server/openapi -type f -not -name '*_service.go' -delete
