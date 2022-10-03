package bootstrap

import (
	"api-template/config"
	"api-template/internal/platform/server/handler/health"
	server "api-template/internal/platform/server/openapi"
	"fmt"
	"net/http"
)

// RunHealth starts a server for healthcheck status
func RunHealth() error {
	addr := fmt.Sprintf(":%d", config.AppConfig.HttpHealthPort)

	http.HandleFunc("/health", health.GetHealth().Handler)

	return http.ListenAndServe(addr, nil)
}

// RunServer serves de API
func RunServer() error {
	addr := fmt.Sprintf(":%d", config.AppConfig.HttpPort)

	// system
	SystemApiService := server.NewSystemApiService()
	SystemApiController := server.NewSystemApiController(SystemApiService)

	router := server.NewRouter(SystemApiController)

	return http.ListenAndServe(addr, router)
}
