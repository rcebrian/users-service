package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/internal/platform/storage/mysql"
	"api-template/pkg/logger"
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

func init() {
	if err := envconfig.Process("", &config.AppConfig); err != nil {
		logger.Fatal("APP environment variables could not be processed")
	}

	if err := envconfig.Process("", &config.ServerConfig); err != nil {
		logger.Fatal("SERVER environment variables could not be processed")
	}

	if err := envconfig.Process("", &config.MySqlConfig); err != nil {
		logger.Fatal("DATABASE environment variables could not be processed")
	}

	if err := logger.ParseLevel(config.AppConfig.LogLevel); err != nil {
		logger.Fatal("error parsing log level")
	}

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	go func() {
		logger.Debugf("healthcheck running on :%d/health", config.AppConfig.HttpInternalPort)

		if err := bootstrap.RunInternalServer(); err != nil {
			logger.Fatal(err)
		}
	}()
}

func main() {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MySqlConfig.User, config.MySqlConfig.Passwd, config.MySqlConfig.Host, config.MySqlConfig.Port, config.MySqlConfig.Database)
	db, _ := sql.Open("mysql", mysqlURI)

	userRepo := mysql.NewUserRepository(db)

	var gracefulTime = time.Second * time.Duration(config.ServerConfig.GracefulTime)

	srv := bootstrap.NewServer(userRepo)

	// https://github.com/gorilla/mux#graceful-shutdown
	go func() {
		logger.Infof("http server starting on port :%d", config.ServerConfig.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), gracefulTime)
	defer cancel()

	_ = srv.Shutdown(ctx)

	logger.Warn("http server closed")
	os.Exit(0)
}
