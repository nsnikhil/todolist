package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/magiconair/properties/assert"
	"testing"
	"todolist/domain"
	"todolist/store"
)

func TestCreateNewTodoListService(t *testing.T) {
	testCases := []struct {
		name            string
		actualService   func() TodoListServiceInterface
		expectedService TodoListService
	}{
		{
			name: "test create new todolist service",
			actualService: func() TodoListServiceInterface {
				appStore := store.NewStore(&sqlx.DB{})
				return NewTodoListService(appStore.GetTodoListStore())
			},
			expectedService: TodoListService{todoListStore: store.NewTodoListStore(&sqlx.DB{})},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedService, testCase.actualService())
		})
	}
}

func TestTodoListServiceAddTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "add task buy biscuits",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Add", mockTask).Return(nil)

				service := NewTodoListService(mockTodoListStore)

				return service.Add(mockTask)
			},
			expectedError: nil,
		},
		{
			name: "add task failed",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Add", mockTask).Return(errors.New("failed to add item"))

				service := NewTodoListService(mockTodoListStore)

				return service.Add(mockTask)
			},
			expectedError: errors.New("failed to add item"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedError, testCase.actualError())
		})
	}
}

func TestTodoListServiceRemoveTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "test remove item",
			actualError: func() error {
				mockTask := &domain.MockTask{}
				id := uuid.New().String()

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Remove", id).Return(nil)
				mockTask.On("GetID").Return(id)

				service := NewTodoListService(mockTodoListStore)

				return service.Remove(mockTask.GetID())
			},
			expectedError: nil,
		},
		{
			name: "test remove item failure",
			actualError: func() error {
				mockTask := &domain.MockTask{}
				id := uuid.New().String()

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Remove", id).Return(errors.New("failed to remove task"))
				mockTask.On("GetID").Return(id)

				service := NewTodoListService(mockTodoListStore)

				return service.Remove(mockTask.GetID())
			},
			expectedError: errors.New("failed to remove task"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedError, testCase.actualError())
		})
	}
}

func TestTodoListServiceUpdateTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "test update task",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Update", mockTask).Return(nil)

				service := NewTodoListService(mockTodoListStore)

				return service.Update(mockTask)
			},
			expectedError: nil,
		},
		{
			name: "test update task failure items",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("Update", mockTask).Return(errors.New("failed to update task"))

				service := NewTodoListService(mockTodoListStore)

				return service.Update(mockTask)
			},
			expectedError: errors.New("failed to update task"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedError, testCase.actualError())
		})
	}
}

func TestTodoListServiceGetTodoList(t *testing.T) {
	testCases := []struct {
		name             string
		actualTodoList   func() (domain.TodoListInterface, error)
		expectedTodoList domain.TodoListInterface
		expectedError    error
	}{
		{
			name: "test get todolist",
			actualTodoList: func() (domain.TodoListInterface, error) {

				mockTodoList := &domain.MockTodoList{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("GetTodoList").Return(mockTodoList, nil)

				service := NewTodoListService(mockTodoListStore)

				return service.GetTodoList()
			},
			expectedTodoList: &domain.MockTodoList{},
		},
		{
			name: "test get todolist failure",
			actualTodoList: func() (domain.TodoListInterface, error) {

				mockTodoList := &domain.MockTodoList{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("GetTodoList").Return(mockTodoList, errors.New("failed to get todolist"))

				service := NewTodoListService(mockTodoListStore)

				return service.GetTodoList()
			},
			expectedError: errors.New("failed to get todolist"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			todoList, err := testCase.actualTodoList()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTodoList, todoList)
		})
	}
}

func TestTodoListServiceGetTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualTask    func() (domain.TaskInterface, error)
		expectedTask  domain.TaskInterface
		expectedError error
	}{
		{
			name: "test get task",
			actualTask: func() (domain.TaskInterface, error) {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("GetTask", "uuid").Return(mockTask, nil)

				service := NewTodoListService(mockTodoListStore)
				return service.GetTask("uuid")
			},
			expectedTask: &domain.MockTask{},
		},
		{
			name: "test get task failure",
			actualTask: func() (domain.TaskInterface, error) {
				mockTask := &domain.MockTask{}

				mockTodoListStore := store.MockTodoListStore{}
				mockTodoListStore.On("GetTask", "uuid").Return(mockTask, errors.New("failed to get task"))

				service := NewTodoListService(mockTodoListStore)
				return service.GetTask("uuid")
			},
			expectedError: errors.New("failed to get task"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task, err := testCase.actualTask()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTask, task)
		})
	}
}
