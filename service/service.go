package service

import (
	"todolist/applogger"
	"todolist/constants"
)

type Service struct {
	todoListService TodoListServiceInterface
}

func NewService(todoListService TodoListServiceInterface) Service {
	applogger.Infof(constants.NewService, "[Service] [NewService]")
	return Service{
		todoListService: todoListService,
	}
}

func (s Service) GetTodoListService() TodoListServiceInterface {
	applogger.Infof(constants.ServiceGetTodoListService, "[Service] [GetTodoListService]")
	return s.todoListService
}
