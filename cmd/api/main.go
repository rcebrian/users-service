package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/internal/platform/storage/mysql"
	"api-template/pkg/yaml"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	var (
		err   error
		level logrus.Level
	)

	if err = envconfig.Process("", &config.ServiceConfig); err != nil {
		logrus.WithError(err).Fatal("APP environment variables could not be processed")
	}

	if err = envconfig.Process("", &config.HttpServerConfig); err != nil {
		logrus.WithError(err).Fatal("SERVER environment variables could not be processed")
	}

	if err = envconfig.Process("", &config.MySqlConfig); err != nil {
		logrus.WithError(err).Fatal("DATABASE environment variables could not be processed")
	}

	if level, err = logrus.ParseLevel(config.ServiceConfig.LogLevel); err != nil {
		logrus.WithError(err).Fatal("error parsing log level")
	}

	logrus.SetLevel(level)

	loadOASpecs()

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s",
		config.MySqlConfig.User, config.MySqlConfig.Passwd,
		config.MySqlConfig.Host, config.MySqlConfig.Port,
		config.MySqlConfig.Database,
		time.Duration(config.MySqlConfig.Timeout)*time.Second)
	db, _ = sql.Open("mysql", mysqlURI)

	// starts the internal service with private endpoints
	go func() {
		logrus.Debugf("healthcheck running on :%d/health", config.ServiceConfig.HttpInternalPort)

		if err := bootstrap.RunInternalServer(db); err != nil {
			logrus.Fatal(err)
		}
	}()
}

func main() {
	userRepo := mysql.NewUserRepository(db)

	var gracefulTime = time.Second * time.Duration(config.HttpServerConfig.GracefulTime)

	srv := bootstrap.NewServer(userRepo)

	// https://github.com/gorilla/mux#graceful-shutdown
	go func() {
		logrus.Infof("http server starting on port :%d", config.HttpServerConfig.Port)

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

	logrus.Warn("http server closed")
	os.Exit(0)
}

// loadOASpecs loads ServiceID and Version from OpenAPI specs file
func loadOASpecs() {
	oa, _ := yaml.ReadOpenAPI("api/openapi-spec/openapi.yaml")
	config.ServiceConfig.ServiceID = oa.Info.ServiceID
	config.ServiceConfig.ServiceVersion = oa.Info.Version
}
