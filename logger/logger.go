package logger

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func NewLogger()  {
	if logger != nil {
		return
	}
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
	})
	return
}



func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Error(args ...interface{}) {
	logger.Error(args ...)
}

func Warn(args ...interface{}) {
	logger.Warn(args ...)
}


func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

