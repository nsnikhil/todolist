package service

import (
	"github.com/stretchr/testify/mock"
	"todolist/domain"
)

type MockTodoListServiceInterface struct {
	mock.Mock
}

func (mock MockTodoListServiceInterface) Add(task domain.TaskInterface) error {
	args := mock.Called(task)
	return args.Error(0)
}

func (mock MockTodoListServiceInterface) Remove(taskID string) error {
	args := mock.Called(taskID)
	return args.Error(0)
}

func (mock MockTodoListServiceInterface) Update(task domain.TaskInterface) error {
	args := mock.Called(task)
	return args.Error(0)
}

func (mock MockTodoListServiceInterface) GetTodoList() (domain.TodoListInterface, error) {
	args := mock.Called()
	return args.Get(0).(domain.TodoListInterface), args.Error(1)
}

func (mock MockTodoListServiceInterface) GetTask(taskID string) (domain.TaskInterface, error) {
	args := mock.Called(taskID)
	return args.Get(0).(domain.TaskInterface), args.Error(1)
}
