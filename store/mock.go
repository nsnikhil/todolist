package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"todolist/domain"
)

type MockDbHandle struct {
	mock.Mock
}

func (mock *MockDbHandle) GetDB() (*sqlx.DB, error) {
	args := mock.Called()
	return args.Get(0).(*sqlx.DB), args.Error(1)
}

type MockTaskStore struct {
	mock.Mock
}

func (mock *MockTaskStore) Add(task domain.Task) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}

func (mock *MockTaskStore) Remove(id string, ids ...string) (int64, error) {
	args := mock.Called(id, ids)
	return args.Get(0).(int64), args.Error(1)
}

func (mock *MockTaskStore) Update(task domain.Task) (int64, error) {
	args := mock.Called(task)
	return args.Get(0).(int64), args.Error(1)
}

func (mock *MockTaskStore) GetTasks(ids ...string) ([]domain.Task, error) {
	args := mock.Called(ids)
	return args.Get(0).([]domain.Task), args.Error(1)
}
