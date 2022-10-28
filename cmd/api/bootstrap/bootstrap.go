package bootstrap

import (
	"api-template/config"
	"api-template/internal/platform/server/handler/health"
	server "api-template/internal/platform/server/openapi"
	"api-template/internal/platform/storage/mysql"
	"database/sql"
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
func NewServer() *http.Server {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MySqlConfig.User, config.MySqlConfig.Passwd, config.MySqlConfig.Host, config.MySqlConfig.Port, config.MySqlConfig.Database)
	db, _ := sql.Open("mysql", mysqlURI)

	addr := fmt.Sprintf(":%d", config.ServerConfig.Port)

	// users
	UsersApiController := usersApiController(db)

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
func usersApiController(db *sql.DB) server.Router {
	userRepo := mysql.NewUserRepository(db)

	UsersApiService := server.NewUsersApiService(userRepo)

	return server.NewUsersApiController(UsersApiService)
}
