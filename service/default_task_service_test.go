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

func TestCreateNewTaskService(t *testing.T) {
	testCases := []struct {
		name            string
		actualService   func() TaskService
		expectedService TaskService
	}{
		{
			name: "test create new todolist service",
			actualService: func() TaskService {
				appStore := store.NewStore(&sqlx.DB{})
				return NewTaskService(appStore.GetTodoListStore())
			},
			expectedService: DefaultTaskService{taskStore: store.NewTaskStore(&sqlx.DB{})},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedService, testCase.actualService())
		})
	}
}

func TestTaskServiceAddTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "add task buy biscuits",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Add", mockTask).Return(nil)

				service := NewTaskService(mockTodoListStore)

				return service.Add(mockTask)
			},
			expectedError: nil,
		},
		{
			name: "add task failed",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Add", mockTask).Return(errors.New("failed to add item"))

				service := NewTaskService(mockTodoListStore)

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

func TestTaskServiceRemoveTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "test remove item",
			actualError: func() error {
				id := uuid.New().String()

				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return(id)

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Remove", id, []string(nil)).Return(nil)

				service := NewTaskService(mockTodoListStore)

				return service.Remove(mockTask.GetID())
			},
			expectedError: nil,
		},
		{
			name: "test remove item failure",
			actualError: func() error {
				id := uuid.New().String()

				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return(id)

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Remove", id, []string(nil)).Return(errors.New("failed to remove task"))

				service := NewTaskService(mockTodoListStore)

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

func TestTaskServiceUpdateTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualError   func() error
		expectedError error
	}{
		{
			name: "test update task",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Update", mockTask).Return(nil)

				service := NewTaskService(mockTodoListStore)

				return service.Update(mockTask)
			},
			expectedError: nil,
		},
		{
			name: "test update task failure items",
			actualError: func() error {
				mockTask := &domain.MockTask{}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("Update", mockTask).Return(errors.New("failed to update task"))

				service := NewTaskService(mockTodoListStore)

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

func TestTaskServiceGetTasks(t *testing.T) {
	testCases := []struct {
		name          string
		actualTask    func() ([]domain.Task, error)
		expectedTask  []domain.Task
		expectedError error
	}{
		{
			name: "test server tasks",
			actualTask: func() ([]domain.Task, error) {
				mockTasks := []domain.Task{
					&domain.MockTask{},
				}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("GetTasks", []string{"uuid"}).Return(mockTasks, nil)

				service := NewTaskService(mockTodoListStore)
				return service.GetTasks("uuid")
			},
			expectedTask: []domain.Task{
				&domain.MockTask{},
			},
		},
		{
			name: "test server task failure",
			actualTask: func() ([]domain.Task, error) {
				mockTasks := []domain.Task{
					&domain.MockTask{},
				}

				mockTodoListStore := &store.MockTaskStore{}
				mockTodoListStore.On("GetTasks", []string{"uuid"}).Return(mockTasks, errors.New("failed to server task"))

				service := NewTaskService(mockTodoListStore)
				return service.GetTasks("uuid")
			},
			expectedError: errors.New("failed to server task"),
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
