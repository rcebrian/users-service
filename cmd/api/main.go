package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/pkg/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := envconfig.Process("", &config.AppConfig); err != nil {
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
	logger.Infof("Server starting on port :%d", config.AppConfig.HttpPort)

	if err := bootstrap.RunServer(); err != nil {
		logger.Fatal(err)
	}

}
