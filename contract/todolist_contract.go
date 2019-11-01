package contract

import (
	"github.com/google/uuid"
	"todolist/applogger"
	"todolist/constants"
)

type TodoListResponse struct {
	Data         interface{} `json:"data,omitempty"`
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"error"`
}

func NewTodoListResponse(data interface{}, errorMessage string, success bool) TodoListResponse {
	applogger.Infof(constants.NewTodoListResponse, "[TodoListResponse] [NewTodoListResponse]")
	return TodoListResponse{
		Data:         data,
		Success:      success,
		ErrorMessage: errorMessage,
	}
}

type TaskIDRequest struct {
	TaskID string `json:"task_id"`
}

func NewTaskIDRequest(taskID string) TaskIDRequest {
	return TaskIDRequest{TaskID: taskID}
}

func (rtr TaskIDRequest) Validate() bool {
	applogger.Infof(constants.TaskIdRequestValidate, "[TaskIDRequest] [Validate]")
	if _, err := uuid.Parse(rtr.TaskID); err != nil {
		return false
	}
	return true
}
