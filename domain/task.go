package domain

import (
	"fmt"
	"github.com/google/uuid"
	"todolist/apperror"
	"todolist/applogger"
	"todolist/constants"
)

type TaskInterface interface {
	ChangeDescription(newDescription string)
	ChangeStatus()
	GetDescription() string
	GetStatus() bool
	GetID() string
	Validate() bool
}

type Task struct {
	Id          string `db:"id"`
	Description string `db:"description"`
	Status      bool   `db:"status"`
}

func NewTask(id string, description string, status bool) (*Task, error) {
	if len(description) == 0 {
		applogger.Errorf(constants.InvalidTaskErrorMessage, fmt.Sprint("[Task] [NewTask]"))
		return nil, apperror.NewInvalidTaskError(constants.InvalidTaskErrorMessage, fmt.Sprint("[Task] [NewTask]"))
	}
	if _, err := uuid.Parse(id); err != nil {
		applogger.Errorf(constants.InvalidTaskUUIDErrorMessage, fmt.Sprint("[Task] [NewTask]"), id)
		return nil, apperror.NewInvalidTaskError(constants.InvalidTaskUUIDErrorMessage, fmt.Sprint("[Task] [NewTask]"), id)
	}
	applogger.Infof(constants.TaskCreationSuccessMessage, fmt.Sprint("[Task] [NewTask]"), description)
	return &Task{Id: id, Description: description, Status: status}, nil
}

func (t *Task) ChangeDescription(newDescription string) {
	oldDescription := t.Description
	t.Description = newDescription
	applogger.Infof(constants.TaskDescriptionChangeSuccessMessage, fmt.Sprint("[Task] [ChangeDescription]"), oldDescription, newDescription)
}

func (t *Task) ChangeStatus() {
	oldStatus := t.Status
	t.Status = !t.Status
	applogger.Infof(constants.TaskStatusChangeSuccessMessage, fmt.Sprint("[Task] [ChangeStatus]"), oldStatus, t.Status)
}

func (t *Task) GetDescription() string {
	applogger.Infof(constants.TaskGetDescriptionLog, fmt.Sprint("[Task] [GetDescription]"), t.Description)
	return t.Description
}

func (t *Task) GetStatus() bool {
	applogger.Infof(constants.TaskGetStatusLog, fmt.Sprint("[Task] [GetStatus]"), t.Status)
	return t.Status
}

func (t *Task) GetID() string {
	applogger.Infof(constants.TaskGetIDLog, fmt.Sprint("[Task] [GetID]"), t.Id)
	return t.Id
}

func (t *Task) Validate() bool {
	if len(t.Description) == 0 {
		applogger.Errorf(constants.InvalidTaskErrorMessage, fmt.Sprint("[Task] [NewTask]"))
		return false
	}

	if _, err := uuid.Parse(t.Id); err != nil {
		applogger.Errorf(constants.InvalidTaskUUIDErrorMessage, fmt.Sprint("[Task] [NewTask]"), t.Id)
		return false
	}

	applogger.Infof(constants.SuccessfulTaskValidation, fmt.Sprint("[Task] [NewTask]"), t)
	return true
}
