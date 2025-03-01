package logger

import (
	"github.com/Crabocod/gpt_network/generate-service/internal/config"

	"github.com/sirupsen/logrus"
)

var Logrus *logrus.Logger

func LoadLogger() error {
	Logrus = logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true
	Logrus.SetFormatter(Formatter)
	level, err := logrus.ParseLevel(config.Data.Logger.LogLevel)
	if err != nil {
		return err
	}

	Logrus.SetLevel(level)
	return nil
}
