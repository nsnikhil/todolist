package service

import (
	"github.com/stretchr/testify/mock"
	"todolist/domain"
)

type MockTaskService struct {
	mock.Mock
}

func (mock *MockTaskService) Add(task domain.Task) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}

func (mock *MockTaskService) Remove(id string, ids ...string) (int64, error) {
	args := mock.Called(id, ids)
	return args.Get(0).(int64), args.Error(1)
}

func (mock *MockTaskService) Update(task domain.Task) (int64, error) {
	args := mock.Called(task)
	return args.Get(0).(int64), args.Error(1)
}

func (mock *MockTaskService) GetTasks(ids ...string) ([]domain.Task, error) {
	args := mock.Called(ids)
	return args.Get(0).([]domain.Task), args.Error(1)
}
