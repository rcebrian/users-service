package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

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

// NewHealthServer starts a server for healthcheck status
func NewHealthServer(sqlClient *sql.DB) *http.Server {
	if err := envconfig.Process("HEALTH", &configs.HealthHttpServerConfig); err != nil {
		logrus.WithError(err).Fatal("HEALTH environment variables could not be processed")
	}

	healthHandler := http.NewServeMux()

	mysqlAffectedEndpoints := []string{"/users", "/user"}

	mysqlHealth := providers.NewMysqlProvider("mysql", mysqlAffectedEndpoints, sqlClient, configs.MySqlConfig.Timeout, configs.MySqlConfig.Threshold)
	healthService := health.NewHealthService(configs.ServiceConfig.ServiceID, configs.ServiceConfig.ServiceVersion, mysqlHealth)

	healthHandler.HandleFunc("/health", healthService.Handler)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", configs.HealthHttpServerConfig.Port),
		Handler:      healthHandler,
		WriteTimeout: configs.HealthHttpServerConfig.WriteTimeout,
		ReadTimeout:  configs.HealthHttpServerConfig.ReadTimeout,
		IdleTimeout:  configs.HealthHttpServerConfig.IdleTimeout,
	}
}

// NewServer create a new configured server
func NewServer(userRepo users.UserRepository) *http.Server {
	router := chi.NewRouter()

	router.Use(middlewares.PanicRecovery)
	router.Use(middlewares.Logging)
	router.Use(middlewares.Cors)

	apiDoc := redoc.Redoc{
		Title:       "API Docs",
		Description: "API documentation",
		SpecFile:    "./api/openapi-specs/openapi.yaml",
		SpecPath:    "/openapi.yaml",
		DocsPath:    "/docs",
	}
	router.Method(http.MethodGet, apiDoc.DocsPath, apiDoc.Handler())
	router.Method(http.MethodGet, apiDoc.SpecPath, apiDoc.Handler())

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
