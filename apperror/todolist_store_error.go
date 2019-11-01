package apperror

import "fmt"

type TaskAddError struct {
	ErrorFormat string
	Arguments   []string
}

func NewTaskAddError(errorFormat string, arguments ...string) TaskAddError {
	return TaskAddError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e TaskAddError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}

type TaskRemoveError struct {
	ErrorFormat string
	Arguments   []string
}

func NewTaskRemoveError(errorFormat string, arguments ...string) TaskRemoveError {
	return TaskRemoveError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e TaskRemoveError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}

type TaskUpdateError struct {
	ErrorFormat string
	Arguments   []string
}

func NewTaskUpdateError(errorFormat string, arguments ...string) TaskUpdateError {
	return TaskUpdateError{ErrorFormat: errorFormat, Arguments: arguments}
}

func (e TaskUpdateError) Error() string {
	return fmt.Sprintf(e.ErrorFormat, e.Arguments)
}
