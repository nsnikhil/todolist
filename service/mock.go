package service

import (
	"github.com/stretchr/testify/mock"
	"todolist/domain"
)

type MockTaskService struct {
	mock.Mock
}

func (mock *MockTaskService) Add(task domain.Task) error {
	args := mock.Called(task)
	return args.Error(0)
}

func (mock *MockTaskService) Remove(id string, ids ...string) error {
	args := mock.Called(id, ids)
	return args.Error(0)
}

func (mock *MockTaskService) Update(task domain.Task) error {
	args := mock.Called(task)
	return args.Error(0)
}

func (mock *MockTaskService) GetTasks(ids ...string) ([]domain.Task, error) {
	args := mock.Called(ids)
	return args.Get(0).([]domain.Task), args.Error(1)
}
