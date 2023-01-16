package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rcebrian/users-service/pkg/health"
	"github.com/rcebrian/users-service/pkg/health/providers"

	"github.com/rcebrian/users-service/configs"
	users "github.com/rcebrian/users-service/internal"
	server "github.com/rcebrian/users-service/internal/platform/server/openapi"
	"github.com/rcebrian/users-service/internal/users/creating"
	"github.com/rcebrian/users-service/internal/users/finding"

	"github.com/mvrilo/go-redoc"
)

// RunInternalServer starts a server for healthcheck status
func RunInternalServer(sqlClient *sql.DB) error {
	mysqlAffectedEndpoints := []string{"/users", "/user"}

	mysqlHealth := providers.NewMysqlProvider("mysql", mysqlAffectedEndpoints, sqlClient, configs.MySqlConfig.Timeout, configs.MySqlConfig.Threshold)
	healthService := health.NewHealthService(configs.ServiceConfig.ServiceID, configs.ServiceConfig.ServiceVersion, mysqlHealth)

	addr := fmt.Sprintf(":%d", configs.ServiceConfig.HttpInternalPort)
	internal := http.NewServeMux()
	internal.HandleFunc("/health", healthService.Handler)

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
	addr := fmt.Sprintf(":%d", configs.HttpServerConfig.Port)

	// users
	UsersApiController := usersApiController(userRepo)

	router := server.NewRouter(UsersApiController)

	return &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: configs.HttpServerConfig.WriteTimeout,
		ReadTimeout:  configs.HttpServerConfig.ReadTimeout,
		IdleTimeout:  configs.HttpServerConfig.IdleTimeout,
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
