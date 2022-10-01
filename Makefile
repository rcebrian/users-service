api_gen:
	sh scripts/openapi.sh

api_clean:
	rm -r internal/platform/server/api \
		internal/platform/server/.openapi-generator \
		internal/platform/server/openapi
