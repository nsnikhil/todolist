package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"todolist/domain"
)

type MockDbHandle struct {
	mock.Mock
}

func (mock MockDbHandle) GetDB() (*sqlx.DB, error) {
	args := mock.Called()
	return args.Get(0).(*sqlx.DB), args.Error(1)
}

type MockTodoListStore struct {
	mock.Mock
}

func (mock MockTodoListStore) Add(task domain.TaskInterface) error {
	args := mock.Called(task)
	return args.Error(0)
}
func (mock MockTodoListStore) Remove(taskID string) error {
	args := mock.Called(taskID)
	return args.Error(0)
}
func (mock MockTodoListStore) Update(task domain.TaskInterface) error {
	args := mock.Called(task)
	return args.Error(0)
}
func (mock MockTodoListStore) GetTodoList() (domain.TodoListInterface, error) {
	args := mock.Called()
	return args.Get(0).(domain.TodoListInterface), args.Error(1)
}
func (mock MockTodoListStore) GetTask(taskID string) (domain.TaskInterface, error) {
	args := mock.Called(taskID)
	return args.Get(0).(domain.TaskInterface), args.Error(1)
}
