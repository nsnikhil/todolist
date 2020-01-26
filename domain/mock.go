package domain

import "github.com/stretchr/testify/mock"

type MockTask struct {
	mock.Mock
}

func (mock *MockTask) UpdateTitle(newTitle string) {
	mock.Called(newTitle)
}

func (mock *MockTask) UpdateDescription(newDescription string) {
	mock.Called(newDescription)
}

func (mock *MockTask) UpdateStatus() {
	mock.Called()
}

func (mock *MockTask) UpdateTags(tags ...string) {
	mock.Called(tags)
}

func (mock *MockTask) GetID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *MockTask) GetTitle() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *MockTask) GetDescription() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *MockTask) GetStatus() bool {
	args := mock.Called()
	return args.Bool(0)
}

func (mock *MockTask) GetTags() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

type MockTodoList struct {
	mock.Mock
}

func (mock *MockTodoList) Add(task *DefaultTask) {
	mock.Called(task)
}

func (mock *MockTodoList) Remove(task *DefaultTask) error {
	args := mock.Called(task)
	return args.Error(0)
}

type MockTaskFactory struct {
	mock.Mock
}

func (mock *MockTaskFactory) Create(title string, description string, status bool, tags ...string) (Task, error) {
	args := mock.Called(title, description, status, tags)
	return args.Get(0).(Task), args.Error(1)
}

func (mock *MockTaskFactory) CreateWithID(id string, title string, description string, status bool, tags ...string) (Task, error) {
	args := mock.Called(id, title, description, status, tags)
	return args.Get(0).(Task), args.Error(1)
}
