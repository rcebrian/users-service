package configs

import "time"

var ServiceConfig ServiceConf

type ServiceConf struct {
	ServiceID        string
	ServiceVersion   string
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	HttpInternalPort int    `envconfig:"HTTP_INTERNAL_PORT" default:"8079"`
}

var HttpServerConfig HttpServerConf

type HttpServerConf struct {
	Port         int           `envconfig:"HTTP_PORT" default:"8080"`
	GracefulTime time.Duration `envconfig:"HTTP_GRACEFUL_TIME" default:"30s"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"10s"`
	IdleTimeout  time.Duration `envconfig:"HTTP_IDLE_TIMEOUT" default:"90s"`
}

var MySqlConfig MySqlConf

type MySqlConf struct {
	Host      string        `envconfig:"MYSQL_HOST" default:"localhost"`
	Port      int           `envconfig:"MYSQL_PORT" default:"3306"`
	Timeout   time.Duration `envconfig:"MYSQL_TIMEOUT" default:"5000ms"`
	Threshold time.Duration `envconfig:"MYSQL_THRESHOLD" default:"500ms"`
	User      string        `envconfig:"MYSQL_USER" default:"srvuser"`
	Passwd    string        `envconfig:"MYSQL_PASSWD" default:"srvuser"`
	Database  string        `envconfig:"MYSQL_DATABASE" default:"users"`
}
