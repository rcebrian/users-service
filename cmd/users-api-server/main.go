package main

import (
	"context"
	"errors"
	apiConfig "github.com/rcebrian/users-service/configs/users-api-server"
	"net/http"
	"os"
	"os/signal"

	"github.com/rcebrian/users-service/cmd/users-api-server/bootstrap"
	"github.com/rcebrian/users-service/configs"
	"github.com/sirupsen/logrus"
)

func init() {
	err := apiConfig.ConfigureServer()
	if err != nil {
		logrus.WithError(err).Fatalf("error configuring %s application", apiConfig.ServiceID)
	}

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
		logrus.Infof("HTTP server starting on port %s", server.Addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.WithError(err).Fatal("error starting HTTP server")
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
