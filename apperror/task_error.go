package apperror

import (
	"fmt"
)

type InvalidTaskError struct {
	ErrorFormat string
	Arguments   []string
}

func NewInvalidTaskError(errorFormat string, arguments ...string) InvalidTaskError {
	return InvalidTaskError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e InvalidTaskError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}

type TaskRemoveFailedError struct {
	ErrorFormat string
	Arguments   []string
}

func NewTaskRemoveFailedError(errorFormat string, arguments ...string) TaskRemoveFailedError {
	return TaskRemoveFailedError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e TaskRemoveFailedError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}
