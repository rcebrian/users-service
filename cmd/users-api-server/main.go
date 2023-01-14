package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rcebrian/users-service/cmd/users-api-server/bootstrap"
	"github.com/rcebrian/users-service/configs"
	"github.com/rcebrian/users-service/internal/platform/storage/mysql"
	"github.com/rcebrian/users-service/pkg/yaml"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	var (
		err   error
		level logrus.Level
	)

	if err = envconfig.Process("", &configs.ServiceConfig); err != nil {
		logrus.WithError(err).Fatal("APP environment variables could not be processed")
	}

	if err = envconfig.Process("", &configs.HttpServerConfig); err != nil {
		logrus.WithError(err).Fatal("SERVER environment variables could not be processed")
	}

	if err = envconfig.Process("", &configs.MySqlConfig); err != nil {
		logrus.WithError(err).Fatal("DATABASE environment variables could not be processed")
	}

	if level, err = logrus.ParseLevel(configs.ServiceConfig.LogLevel); err != nil {
		logrus.WithError(err).Fatal("error parsing log level")
	}

	logrus.SetLevel(level)

	loadOASpecs()

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s",
		configs.MySqlConfig.User, configs.MySqlConfig.Passwd,
		configs.MySqlConfig.Host, configs.MySqlConfig.Port,
		configs.MySqlConfig.Database,
		time.Duration(configs.MySqlConfig.Timeout)*time.Second)
	db, _ = sql.Open("mysql", mysqlURI)

	// starts the internal service with private endpoints
	go func() {
		logrus.Debugf("healthcheck running on :%d/health", configs.ServiceConfig.HttpInternalPort)

		if err := bootstrap.RunInternalServer(db); err != nil {
			logrus.Fatal(err)
		}
	}()
}

func main() {
	userRepo := mysql.NewUserRepository(db)

	var gracefulTime = time.Second * time.Duration(configs.HttpServerConfig.GracefulTime)

	srv := bootstrap.NewServer(userRepo)

	// https://github.com/gorilla/mux#graceful-shutdown
	go func() {
		logrus.Infof("http server starting on port :%d", configs.HttpServerConfig.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), gracefulTime)
	defer cancel()

	_ = srv.Shutdown(ctx)

	logrus.Info("shutting down HTTP server...")
	os.Exit(0)
}

// loadOASpecs loads ServiceID and Version from OpenAPI specs file
func loadOASpecs() {
	oa, _ := yaml.ReadOpenAPI("api/openapi-spec/openapi.yaml")
	configs.ServiceConfig.ServiceID = oa.Info.ServiceID
	configs.ServiceConfig.ServiceVersion = oa.Info.Version
}
