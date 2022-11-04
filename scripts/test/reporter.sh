#!/usr/bin/env bash
: <<DOCS
Run tests and generate an html report
DOCS

NC='\033[0m'
RED='\033[0;31m'

OUT_DIR="coverage"
OUT_FILE="$OUT_DIR/coverage.out"
REPORT_FILE="$OUT_DIR/coverage.html"

# create coverage directory if not exists
if [ ! -d $OUT_DIR ]; then
  mkdir $OUT_DIR
fi

# run tests
TEST_CMD=$(go test -v -coverprofile $OUT_FILE ./... 1> /dev/null)

if $TEST_CMD; then
  go tool cover -html $OUT_FILE -o $REPORT_FILE
  rm $OUT_FILE
  exit 0
else
  echo -e "${RED}Something gone wrong running html reporter${NC}"
  exit 1
fi

