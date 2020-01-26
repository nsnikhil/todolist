package service

import (
	"todolist/domain"
	"todolist/store"
	"todolist/util"
)

type DefaultTaskService struct {
	taskStore store.TaskStore
}

func NewTaskService(taskStore store.TaskStore) TaskService {
	util.DebugLog("[TaskService] [NewTaskService]", taskStore)
	return DefaultTaskService{taskStore: taskStore}
}

func (tls DefaultTaskService) Add(task domain.Task) error {
	if err := tls.taskStore.Add(task); err != nil {
		return util.LogAndGetError("[DefaultTaskService] [Add]", err)
	}
	util.DebugLog("[DefaultTaskService] [Add]", task)
	return nil
}

func (tls DefaultTaskService) Remove(id string, ids ...string) error {
	if err := tls.taskStore.Remove(id, ids...); err != nil {
		return util.LogAndGetError("[DefaultTaskService] [Remove]", err)
	}
	util.DebugLog("[DefaultTaskService] [Remove]", ids)
	return nil
}

func (tls DefaultTaskService) Update(task domain.Task) error {
	if err := tls.taskStore.Update(task); err != nil {
		return util.LogAndGetError("[DefaultTaskService] [Update]", err)
	}
	util.DebugLog("[DefaultTaskService] [Update]", task)
	return nil
}

func (tls DefaultTaskService) GetTasks(ids ...string) ([]domain.Task, error) {
	task, err := tls.taskStore.GetTasks(ids...)
	if err != nil {
		return nil, util.LogAndGetError("[DefaultTaskService] [GetTasks]", err)
	}
	util.DebugLog("[DefaultTaskService] [GetTasks]", ids)
	return task, err
}
