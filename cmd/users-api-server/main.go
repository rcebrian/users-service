package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	joonix "github.com/joonix/log"

	"github.com/kelseyhightower/envconfig"
	"github.com/rcebrian/users-service/cmd/users-api-server/bootstrap"
	"github.com/rcebrian/users-service/configs"
	"github.com/sirupsen/logrus"
)

func init() {
	var (
		err   error
		level logrus.Level
	)

	logrus.SetFormatter(joonix.NewFormatter(joonix.StackdriverFormat))

	if err = envconfig.Process("", &configs.ServiceConfig); err != nil {
		logrus.WithError(err).Fatal("APP environment variables could not be processed")
	}

	if level, err = logrus.ParseLevel(configs.ServiceConfig.LogLevel); err != nil {
		logrus.WithError(err).Fatal("error parsing log level")
	}

	logrus.SetLevel(level)

	healthServer := bootstrap.NewHealthServer()

	go func() {
		logrus.Infof("healthcheck running on :%d/health", configs.HealthHttpServerConfig.Port)

		if err = healthServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err)
		}
	}()
}

func main() {
	server := bootstrap.NewApiServer()

	go func() {
		logrus.Infof("http server starting on port :%d", configs.HttpServerConfig.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), configs.HttpServerConfig.GracefulTime)
	defer cancel()

	_ = server.Shutdown(ctx)

	logrus.Info("shutting down HTTP server...")
	os.Exit(0)
}
