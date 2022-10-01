package config

const (
	AppName = "api-template"
)

var AppConfig AppConf

type AppConf struct {
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	AppVersion string `envconfig:"APP_VERSION" default:"v1.0.0"`
	HttpPort   int    `envconfig:"HTTP_PORT" default:"8080" `
}
