package zlog

import (
	"github.com/google/martian/v3/log"
	"github.com/sirupsen/logrus"
	"os"
)

var instance = logrus.New()

const defaultLogLevel = "debug"

func init() {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	formatter.ForceColors = true
	instance.SetFormatter(formatter)
	instance.SetOutput(os.Stdout)

	switch defaultLogLevel {
	case "debug":
		instance.SetLevel(logrus.DebugLevel)
	case "warn":
		instance.SetLevel(logrus.WarnLevel)
	case "info":
		instance.SetLevel(logrus.InfoLevel)
	case "error":
		instance.SetLevel(logrus.ErrorLevel)
	default:
		log.Errorf("log level is invalid. loglevel=%s", defaultLogLevel)
	}

	instance.Info("zlog init complete.")
}

func Instance() *logrus.Logger {
	return instance
}
