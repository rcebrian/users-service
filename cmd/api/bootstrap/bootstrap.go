package bootstrap

import (
	"api-template/config"
	server "api-template/internal/platform/server/openapi"
	"fmt"
	"net/http"
)

func Run() error {
	port := fmt.Sprintf(":%d", config.AppConfig.HttpPort)

	// system
	SystemApiService := server.NewSystemApiService()
	SystemApiController := server.NewSystemApiController(SystemApiService)

	router := server.NewRouter(SystemApiController)

	return http.ListenAndServe(port, router)
}
