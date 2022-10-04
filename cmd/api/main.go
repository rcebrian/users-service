package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/pkg/logger"
	"context"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func init() {
	if err := envconfig.Process("", &config.AppConfig); err != nil {
		logger.Fatal("APP environment variables could not be processed")
	}

	if err := envconfig.Process("", &config.ServerConfig); err != nil {
		logger.Fatal("environment variables could not be processed")
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
	var wait time.Duration
	srv := bootstrap.NewServer()

	// https://github.com/gorilla/mux#graceful-shutdown
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Infof("server starting on port :%d", config.ServerConfig.Port)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Warn("server shutting down")
	os.Exit(0)
}
