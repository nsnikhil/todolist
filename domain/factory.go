package domain

import (
	"fmt"
	"github.com/google/uuid"
	"todolist/util"
)

type TaskFactory interface {
	Create(title string, description string, status bool, tags ...string) (Task, error)
	CreateWithID(id string, title string, description string, status bool, tags ...string) (Task, error)
}

type DefaultTaskFactory struct {
}

func NewTaskFactory() TaskFactory {
	util.DebugLog("[TaskFactory] [NewTaskFactory]")
	return DefaultTaskFactory{}
}

func (dtf DefaultTaskFactory) Create(title string, description string, status bool, tags ...string) (Task, error) {
	if err := validate(title); err != nil {
		return nil, err
	}
	util.DebugLog("[DefaultTaskFactory] [Create]")
	return createDefaultTask(uuid.New().String(), title, description, status, tags...), nil
}

func (dtf DefaultTaskFactory) CreateWithID(id string, title string, description string, status bool, tags ...string) (Task, error) {
	if err := validateUUID(id); err != nil {
		return nil, err
	}

	if err := validate(title); err != nil {
		return nil, err
	}

	util.DebugLog("[DefaultTaskFactory] [CreateWithID]")
	return createDefaultTask(id, title, description, status, tags...), nil
}

func createDefaultTask(id string, title string, description string, status bool, tags ...string) *DefaultTask {
	return &DefaultTask{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Tags:        tags,
	}
}

func validateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return util.LogAndGetError("[DefaultTaskFactory] [validateUUID]", err)
	}
	return nil
}

func validate(title string) error {
	if len(title) == 0 {
		return util.LogAndGetError("[DefaultTaskFactory] [validate]", fmt.Errorf("title cannot be empty"))
	}
	return nil
}
