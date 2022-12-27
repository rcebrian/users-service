package configs

var ServiceConfig ServiceConf

type ServiceConf struct {
	ServiceID        string
	ServiceVersion   string
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	HttpInternalPort int    `envconfig:"HTTP_INTERNAL_PORT" default:"8079"`
}

var HttpServerConfig HttpServerConf

type HttpServerConf struct {
	Port         int `envconfig:"HTTP_PORT" default:"8080"`
	GracefulTime int `envconfig:"HTTP_GRACEFUL_TIME" default:"30"`
	WriteTimeout int `envconfig:"HTTP_WRITE_TIMEOUT" default:"15"`
	ReadTimeout  int `envconfig:"HTTP_READ_TIMEOUT" default:"15"`
	IdleTimeout  int `envconfig:"HTTP_IDLE_TIMEOUT" default:"60"`
}

var MySqlConfig MySqlConf

type MySqlConf struct {
	Host      string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port      int    `envconfig:"MYSQL_PORT" default:"3306"`
	Timeout   int    `envconfig:"MYSQL_TIMEOUT" default:"10"`
	Threshold int    `envconfig:"MYSQL_TIMEOUT" default:"30"`
	User      string `envconfig:"MYSQL_USER" default:"srvuser"`
	Passwd    string `envconfig:"MYSQL_PASSWD" default:"srvuser"`
	Database  string `envconfig:"MYSQL_DATABASE" default:"users"`
}
