package logger

import "github.com/sirupsen/logrus"

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
