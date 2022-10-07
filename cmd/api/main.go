package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/pkg/logger"
	"context"
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

	if err := logger.ParseLevel(config.AppConfig.LogLevel); err != nil {
		logger.Fatal("error parsing log level")
	}

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	go func() {
		logger.Debugf("healthcheck running on :%d/health", config.AppConfig.HttpHealthPort)

		if err := bootstrap.RunHealth(); err != nil {
			logger.Fatal(err)
		}
	}()
}

func main() {
	var gracefulTime = time.Second * time.Duration(config.ServerConfig.GracefulTime)

	srv := bootstrap.NewServer()

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
