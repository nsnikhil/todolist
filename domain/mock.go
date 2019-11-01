package domain

import "github.com/stretchr/testify/mock"

type MockTask struct {
	mock.Mock
}

func (mock *MockTask) ChangeDescription(newDescription string) {
	mock.Called(newDescription)
}

func (mock *MockTask) ChangeStatus() {
	mock.Called()
}

func (mock *MockTask) GetDescription() string {
	args := mock.Called()
	return args.String(0)
}
func (mock *MockTask) GetStatus() bool {
	args := mock.Called()
	return args.Bool(0)
}

func (mock *MockTask) GetID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *MockTask) Validate() bool {
	args := mock.Called()
	return args.Bool(0)
}

type MockTodoList struct {
	mock.Mock
}

func (mock *MockTodoList) Add(task *Task) {
	mock.Called(task)
}

func (mock *MockTodoList) Remove(task *Task) error {
	args := mock.Called(task)
	return args.Error(0)
}
