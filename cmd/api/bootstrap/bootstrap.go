package bootstrap

import (
	"api-template/config"
	"api-template/internal/platform/server/handler/health"
	server "api-template/internal/platform/server/openapi"
	"fmt"
	"net/http"
	"time"
)

// RunHealth starts a server for healthcheck status
func RunHealth() error {
	addr := fmt.Sprintf(":%d", config.AppConfig.HttpHealthPort)

	http.HandleFunc("/health", health.GetHealth().Handler)

	return http.ListenAndServe(addr, nil)
}

// NewServer create a new configured server
func NewServer() *http.Server {
	addr := fmt.Sprintf(":%d", config.ServerConfig.Port)

	// system
	SystemApiService := server.NewSystemApiService()
	SystemApiController := server.NewSystemApiController(SystemApiService)

	router := server.NewRouter(SystemApiController)

	return &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(config.ServerConfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.ServerConfig.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.ServerConfig.IdleTimeout),
		Handler:      router,
	}
}
