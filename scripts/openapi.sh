#!/usr/bin/env bash

SPECS_FILE="api/openapi-spec/openapi.yaml"
CONFIG_FILE="api/openapi-spec/config.yaml"

# swagger server generate docs: https://goswagger.io/generate/server.html

openapi-generator-cli generate --generator-name go-server \
  -i $SPECS_FILE \
  -c $CONFIG_FILE \
  --global-property apiDocs=true \
  --global-property verbose=false \
  -o internal/platform/server
