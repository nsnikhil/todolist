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

func (tls DefaultTaskService) Add(task domain.Task) (string, error) {
	id, err := tls.taskStore.Add(task)
	if err != nil {
		return id, util.LogAndGetError("[DefaultTaskService] [Add]", err)
	}
	util.DebugLog("[DefaultTaskService] [Add]", task)
	return id, err
}

func (tls DefaultTaskService) Remove(id string, ids ...string) (int64, error) {
	count, err := tls.taskStore.Remove(id, ids...)
	if err != nil {
		return count, util.LogAndGetError("[DefaultTaskService] [Remove]", err)
	}
	util.DebugLog("[DefaultTaskService] [Remove]", ids)
	return count, err
}

func (tls DefaultTaskService) Update(task domain.Task) (int64, error) {
	count, err := tls.taskStore.Update(task)
	if err != nil {
		return count, util.LogAndGetError("[DefaultTaskService] [Update]", err)
	}
	util.DebugLog("[DefaultTaskService] [Update]", task)
	return count, nil
}

func (tls DefaultTaskService) GetTasks(ids ...string) ([]domain.Task, error) {
	task, err := tls.taskStore.GetTasks(ids...)
	if err != nil {
		return nil, util.LogAndGetError("[DefaultTaskService] [GetTasks]", err)
	}
	util.DebugLog("[DefaultTaskService] [GetTasks]", ids)
	return task, err
}
