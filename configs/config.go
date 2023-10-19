package configs

import (
	"fmt"
	joonix "github.com/joonix/log"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"time"
)

type envLogConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

func ConfigureLogger() error {
	logrus.SetFormatter(joonix.NewFormatter(joonix.StackdriverFormat))

	var (
		err       error
		level     logrus.Level
		logConfig envLogConfig
	)
	if err = envconfig.Process("", &logConfig); err != nil {
		return err
	}

	if level, err = logrus.ParseLevel(logConfig.LogLevel); err != nil {
		return err
	}

	logrus.SetLevel(level)

	return nil
}

var HttpServerConfig HttpServerConf
var HealthHttpServerConfig HttpServerConf

type HttpServerConf struct {
	Port         uint16        `envconfig:"HTTP_PORT" default:"8080"`
	GracefulTime time.Duration `envconfig:"HTTP_GRACEFUL_TIME" default:"30s"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"10s"`
	IdleTimeout  time.Duration `envconfig:"HTTP_IDLE_TIMEOUT" default:"90s"`
}

var MySqlConfig MySqlConf

type MySqlConf struct {
	Host      string        `envconfig:"MYSQL_HOST" default:"localhost"`
	Port      uint16        `envconfig:"MYSQL_PORT" default:"3306"`
	Timeout   time.Duration `envconfig:"MYSQL_TIMEOUT" default:"5000ms"`
	Threshold time.Duration `envconfig:"MYSQL_THRESHOLD" default:"500ms"`
	User      string        `envconfig:"MYSQL_USER" default:"srvuser"`
	Passwd    string        `envconfig:"MYSQL_PASSWD" default:"srvuser"`
	Database  string        `envconfig:"MYSQL_DATABASE" default:"users"`
}

func (mysql MySqlConf) URI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s",
		mysql.User, mysql.Passwd,
		mysql.Host, mysql.Port,
		mysql.Database, mysql.Timeout)
}
