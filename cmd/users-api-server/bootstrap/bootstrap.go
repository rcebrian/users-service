package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rcebrian/users-service/configs"
	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/platform/server"
	"github.com/rcebrian/users-service/internal/users/creating"
	"github.com/rcebrian/users-service/internal/users/finding"
	"github.com/rcebrian/users-service/pkg/health"
	"github.com/rcebrian/users-service/pkg/health/providers"
	"github.com/rcebrian/users-service/pkg/http/server/middlewares"

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
		SpecFile:    "./api/openapi-specs/openapi.yaml",
		SpecPath:    "/openapi.yaml",
		DocsPath:    "/docs",
	}

	internal.HandleFunc(doc.DocsPath, doc.Handler())
	internal.HandleFunc(doc.SpecPath, doc.Handler())

	return http.ListenAndServe(addr, internal)
}

// NewServer create a new configured server
func NewServer(userRepo users.UserRepository) *http.Server {
	router := chi.NewRouter()

	router.Use(middlewares.PanicRecovery)
	router.Use(middlewares.Logging)
	router.Use(middlewares.Cors)

	usersStrictHandler := newApiHandler(userRepo)
	server.HandlerFromMux(usersStrictHandler, router)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", configs.HttpServerConfig.Port),
		Handler:      router,
		WriteTimeout: configs.HttpServerConfig.WriteTimeout,
		ReadTimeout:  configs.HttpServerConfig.ReadTimeout,
		IdleTimeout:  configs.HttpServerConfig.IdleTimeout,
	}
}

// newApiHandler configure users controller with dependency injection
func newApiHandler(userRepo users.UserRepository) server.ServerInterface {
	createService := creating.NewCreatingService(userRepo)
	findAllService := finding.NewFindAllUsersUseCase(userRepo)
	findByIdService := finding.NewFindUserByIdUseCase(userRepo)

	strictServer := server.NewUsersApiServer(createService, findAllService, findByIdService)

	return server.NewStrictHandler(strictServer, nil)
}
