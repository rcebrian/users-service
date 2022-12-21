#!/usr/bin/env bash
: <<DOCS
Generate http server controllers and models via OpenAPI Generator with custom configuration.
DOCS


SPECS_FILE="api/openapi-spec/openapi.yaml"
OUTPUT_DIR="internal/platform/server"

TEMPLATE_ENGINE="mustache"
TEMPLATE_DIR="api/openapi-spec/template"

export GO_POST_PROCESS_FILE="goimports -w"
openapi-generator-cli generate --generator-name go-server \
  --input-spec $SPECS_FILE \
  --template-dir $TEMPLATE_DIR --engine $TEMPLATE_ENGINE \
  --global-property apiDocs=true \
  --global-property verbose=false \
  --enable-post-process-file \
  -p addResponseHeaders=true,featureCORS=true,serverPort=8080,sourceFolder=openapi,packageName=server,outputAsLibrary=true \
  -o $OUTPUT_DIR

if [ $? -eq 0 ]; then
  rm -r ${OUTPUT_DIR}/{.openapi-generator,api}
fi
