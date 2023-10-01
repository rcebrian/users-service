package bootstrap

import (
	"database/sql"
	"fmt"
	users_service "github.com/rcebrian/users-service"
	users_api_config "github.com/rcebrian/users-service/configs/users-api-server"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/mvrilo/go-redoc"
	"github.com/sirupsen/logrus"

	"github.com/rcebrian/users-service/configs"
	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/platform/server"
	"github.com/rcebrian/users-service/internal/platform/storage/mysql"
	"github.com/rcebrian/users-service/internal/users/creating"
	"github.com/rcebrian/users-service/internal/users/finding"

	"github.com/rcebrian/users-service/pkg/health"
	"github.com/rcebrian/users-service/pkg/health/providers"
	"github.com/rcebrian/users-service/pkg/http/server/middlewares"
)

var db *sql.DB

func init() {
	if err := envconfig.Process("", &configs.MySqlConfig); err != nil {
		logrus.WithError(err).Fatal("DATABASE environment variables could not be processed")
	}

	db, _ = sql.Open("mysql", configs.MySqlConfig.URI())
}

// NewHealthServer starts a server for healthcheck status
func NewHealthServer() *http.Server {
	if err := envconfig.Process("HEALTH", &configs.HealthHttpServerConfig); err != nil {
		logrus.WithError(err).Fatal("HEALTH environment variables could not be processed")
	}

	healthHandler := http.NewServeMux()

	mysqlAffectedEndpoints := []string{"/users", "/user"}

	mysqlHealth := providers.NewMysqlProvider("mysql", mysqlAffectedEndpoints, db, configs.MySqlConfig.Timeout, configs.MySqlConfig.Threshold)
	healthService := health.NewHealthService(users_api_config.ServiceID, users_service.VERSION, mysqlHealth)

	healthHandler.HandleFunc("/health", healthService.Handler)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", configs.HealthHttpServerConfig.Port),
		Handler:      healthHandler,
		WriteTimeout: configs.HealthHttpServerConfig.WriteTimeout,
		ReadTimeout:  configs.HealthHttpServerConfig.ReadTimeout,
		IdleTimeout:  configs.HealthHttpServerConfig.IdleTimeout,
	}
}

// NewApiServer create a new configured server
func NewApiServer() *http.Server {
	if err := envconfig.Process("", &configs.HttpServerConfig); err != nil {
		logrus.WithError(err).Fatal("SERVER environment variables could not be processed")
	}

	userRepo := mysql.NewUserRepository(db)

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
