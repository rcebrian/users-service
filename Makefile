SHELL = /bin/bash

api-clean:
	scripts/api/clean.sh

api-codegen: api-clean
	scripts/api/codegen.sh
