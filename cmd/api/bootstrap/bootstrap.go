package bootstrap

import (
	"api-template/config"
	users "api-template/internal"
	"api-template/internal/platform/server/handler/health"
	server "api-template/internal/platform/server/openapi"
	creating "api-template/internal/users"
	"api-template/internal/users/finding"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mvrilo/go-redoc"
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
func NewServer(userRepo users.UserRepository) *http.Server {
	addr := fmt.Sprintf(":%d", config.ServerConfig.Port)

	// users
	UsersApiController := usersApiController(userRepo)

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
func usersApiController(userRepo users.UserRepository) server.Router {
	createService := creating.NewCreatingService(userRepo)
	findAllService := finding.NewFindAllUsersUseCase(userRepo)
	findByIdService := finding.NewFindUserByIdUseCase(userRepo)

	UsersApiService := server.NewUsersApiService(createService, findAllService, findByIdService)

	return server.NewUsersApiController(UsersApiService)
}
