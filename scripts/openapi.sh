#!/usr/bin/env bash

SPECS_FILE="api/openapi-spec/openapi.yaml"
CONFIG_FILE="api/openapi-spec/config.yaml"
OUTPUT_DIR="internal/platform/server"

TEMPLATE_ENGINE="mustache"
TEMPLATE_DIR="api/openapi-spec/template"

export GO_POST_PROCESS_FILE="gofmt -w"
openapi-generator-cli generate --generator-name go-server \
  --input-spec $SPECS_FILE \
  --config $CONFIG_FILE \
  --template-dir $TEMPLATE_DIR --engine $TEMPLATE_ENGINE \
  --global-property apiDocs=true \
  --global-property verbose=false \
  --enable-post-process-file \
  -o $OUTPUT_DIR


# bug: remove empty dir
rmdir $OUTPUT_DIR/api
