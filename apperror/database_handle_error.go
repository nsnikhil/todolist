package apperror

import (
	"fmt"
)

type DatabaseLoadError struct {
	ErrorFormat string
	Arguments   []string
}

func NewDatabaseLoadError(errorFormat string, arguments ...string) DatabaseLoadError {
	return DatabaseLoadError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e DatabaseLoadError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}

type DatabasePingError struct {
	ErrorFormat string
	Arguments   []string
}

func NewDatabasePingError(errorFormat string, arguments ...string) DatabasePingError {
	return DatabasePingError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e DatabasePingError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}
