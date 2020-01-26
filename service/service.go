package service

import (
	"todolist/util"
)

type Service struct {
	TaskService
}

func NewService(taskService TaskService) Service {
	util.DebugLog("[Service] [NewService]")
	return Service{
		TaskService: taskService,
	}
}

func (s Service) GetTaskService() TaskService {
	util.DebugLog("[Service] [GetTaskService]")
	return s.TaskService
}
