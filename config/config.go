package config

var AppConfig AppConf

type AppConf struct {
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	AppVersion       string `envconfig:"APP_VERSION" default:"v1.0.0"`
	HttpInternalPort int    `envconfig:"HTTP_INTERNAL_PORT" default:"8079"`
}

var ServerConfig ServerConf

type ServerConf struct {
	Port         int `envconfig:"HTTP_PORT" default:"8080"`
	GracefulTime int `envconfig:"GRACEFUL_TIME" default:"30"`
	WriteTimeout int `envconfig:"HTTP_WRITE_TIMEOUT" default:"15"`
	ReadTimeout  int `envconfig:"HTTP_READ_TIMEOUT" default:"15"`
	IdleTimeout  int `envconfig:"HTTP_IDLE_TIMEOUT" default:"60"`
}

var MySqlConfig MySqlConf

type MySqlConf struct {
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     int    `envconfig:"MYSQL_PORT" default:"3306"`
	User     string `envconfig:"MYSQL_USER" default:"codely"`
	Passwd   string `envconfig:"MYSQL_PASSWD" default:"codely"`
	Database string `envconfig:"MYSQL_DATABASE" default:"codely"`
}
