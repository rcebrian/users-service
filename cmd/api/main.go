package main

import (
	"api-template/cmd/api/bootstrap"
	"api-template/config"
	"api-template/pkg/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"log"
)

func init() {
	err := envconfig.Process("", &config.AppConfig)
	if err != nil {
		logger.Fatal("environment variables could not be processed")
	}

	_, err = logrus.ParseLevel(config.AppConfig.LogLevel)
	if err != nil {
		logger.Fatal("error parsing log level")
	}

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

}

func main() {
	logger.Infof("Server starting...")

	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
