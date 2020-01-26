package applogger

import (
	"github.com/sirupsen/logrus"
	"os"
	"todolist/config"
)

var appLogger *logrus.Logger

func setupLogger(logLevel string, logFormat string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.WarnLevel
	}

	appLogger = &logrus.Logger{
		Out:       os.Stderr,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}

	if logFormat != "json" {
		appLogger.Formatter = &logrus.TextFormatter{}
	}
}

func init() {
	loadConfig()
	logConfig := config.GetLogConfig()
	setupLogger(logConfig.GetLogLevel(), logConfig.GetLogFormat())
}

func loadConfig() {
	if err := config.Load(); err != nil {
		logrus.Errorf("%s : %v", "[setupDBConnection] [Load]", err)
		panic(err)
	}
}
