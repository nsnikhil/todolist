package util

import "todolist/applogger"

func LogAndGetError(meta string, err error) error {
	LogError(meta, err)
	return err
}

func LogError(meta string, err error) {
	applogger.Errorf("%s : %v", meta, err)
}

func DebugLog(meta string, obj ...interface{}) {
	if obj == nil {
		applogger.Debugf("%s", meta)
		return
	}
	applogger.Debugf("%s : %v", meta, obj)
}
