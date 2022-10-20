package bootstrap

import (
	"api-template/config"
	"api-template/internal/platform/server/handler/health"
	server "api-template/internal/platform/server/openapi"
	"api-template/internal/platform/storage/memory"
	"fmt"
	"net/http"
	"time"

	"github.com/mvrilo/go-redoc"

	"github.com/gorilla/mux"
)

// RunInternalServer starts a server for healthcheck status
func RunInternalServer() error {
	addr := fmt.Sprintf(":%d", config.AppConfig.HttpInternalPort)
	internal := mux.NewRouter()
	internal.HandleFunc("/health", health.GetHealth().Handler)

	doc := redoc.Redoc{
		Title:       "API Docs",
		Description: "API documentation",
		SpecFile:    "./api/openapi-spec/openapi.yaml",
		SpecPath:    "/openapi.yaml",
		DocsPath:    "/docs",
	}

	internal.HandleFunc(doc.DocsPath, doc.Handler())
	internal.HandleFunc(doc.SpecPath, doc.Handler())

	return http.ListenAndServe(addr, internal)
}

// NewServer create a new configured server
func NewServer() *http.Server {
	addr := fmt.Sprintf(":%d", config.ServerConfig.Port)

	// users
	UsersApiController := usersApiController()

	router := server.NewRouter(UsersApiController)

	return &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(config.ServerConfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.ServerConfig.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.ServerConfig.IdleTimeout),
		Handler:      router,
	}
}

// usersApiController configure users controller with dependency injection
func usersApiController() server.Router {
	userRepo := memory.NewUserRepository()

	UsersApiService := server.NewUsersApiService(userRepo)

	return server.NewUsersApiController(UsersApiService)
}
