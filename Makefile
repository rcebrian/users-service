api_gen:
	bash scripts/openapi.sh

api_clean:
	find internal/platform/server/openapi -type f -not -name '*_service.go' -delete
