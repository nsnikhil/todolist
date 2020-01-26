package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualResult  func() (Task, error)
		expectedTask  Task
		expectedError error
	}{
		{
			name: "test create new task",
			actualResult: func() (Task, error) {
				factory := NewTaskFactory()
				return factory.Create("title", "some description", true, "one", "two")
			},
			expectedTask: &DefaultTask{
				Title:       "title",
				Description: "some description",
				Status:      true,
				Tags:        pq.StringArray{"one", "two"},
			},
		},
		{
			name: "test failed to create task because of empty title",
			actualResult: func() (Task, error) {
				factory := NewTaskFactory()
				return factory.Create("", "some description", false)
			},
			expectedError: errors.New("title cannot be empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			if err == nil {
				assert.Equal(t, testCase.expectedTask.GetTitle(), task.GetTitle())
				assert.Equal(t, testCase.expectedTask.GetDescription(), task.GetDescription())
				assert.Equal(t, testCase.expectedTask.GetStatus(), task.GetStatus())
				assert.Equal(t, testCase.expectedTask.GetTags(), task.GetTags())
			}
		})
	}
}

func TestCreateTaskWithID(t *testing.T) {
	id := uuid.New().String()
	testCases := []struct {
		name          string
		actualResult  func() (Task, error)
		expectedTask  Task
		expectedError error
	}{
		{
			name: "test create new task",
			actualResult: func() (Task, error) {
				factory := NewTaskFactory()
				return factory.CreateWithID(id, "title", "some description", true, "one", "two")
			},
			expectedTask: &DefaultTask{
				Id:          id,
				Title:       "title",
				Description: "some description",
				Status:      true,
				Tags:        pq.StringArray{"one", "two"},
			},
		},
		{
			name: "test failed to create task because of invalid uuid",
			actualResult: func() (Task, error) {
				factory := NewTaskFactory()
				return factory.CreateWithID("invalid uuid", "some title", "some description", false)
			},
			expectedError: errors.New("invalid UUID length: 12"),
		},
		{
			name: "test failed to create task because of empty title",
			actualResult: func() (Task, error) {
				factory := NewTaskFactory()
				return factory.CreateWithID(id, "", "some description", false)
			},
			expectedError: errors.New("title cannot be empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			if err == nil {
				assert.Equal(t, testCase.expectedTask.GetTitle(), task.GetTitle())
				assert.Equal(t, testCase.expectedTask.GetDescription(), task.GetDescription())
				assert.Equal(t, testCase.expectedTask.GetStatus(), task.GetStatus())
				assert.Equal(t, testCase.expectedTask.GetTags(), task.GetTags())
			}
		})
	}
}
