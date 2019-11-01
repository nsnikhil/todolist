package applogger

func Tracef(format string, args ...interface{}) {
	appLogger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	appLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	appLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	appLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	appLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	appLogger.Fatalf(format, args...)
}
