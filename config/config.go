package config

var AppConfig AppConf

type AppConf struct {
	LogLevel       string `envconfig:"LOG_LEVEL" default:"info"`
	AppVersion     string `envconfig:"APP_VERSION" default:"v1.0.0"`
	HttpHealthPort int    `envconfig:"HTTP_HEALTH_PORT" default:"8079"`
	HttpPort       int    `envconfig:"HTTP_PORT" default:"8080"`
}
