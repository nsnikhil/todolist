package config

type LogConfig struct {
	logLevel  string
	logFormat string
}

func newLogConfig() LogConfig {
	return LogConfig{
		logLevel:  viperString("LOG_LEVEL"),
		logFormat: viperString("LOG_FORMAT"),
	}
}

func (lc LogConfig) GetLogLevel() string {
	return lc.logLevel
}

func (lc LogConfig) GetLogFormat() string {
	return lc.logFormat
}
