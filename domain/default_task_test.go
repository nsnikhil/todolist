package domain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewTask(t *testing.T) {
	testCases := []struct {
		name          string
		actualTask    func() (Task, error)
		expectedTask  Task
		expectedError error
	}{
		{
			name: "test create new task",
			actualTask: func() (t Task, e error) {
				return NewTaskFactory().Create("buy groceries", "", false)
			},
			expectedTask: &DefaultTask{Title: "buy groceries", Status: false},
		},
		{
			name: "test create completed task",
			actualTask: func() (t Task, e error) {
				return NewTaskFactory().Create("renew id card", "", true)
			},
			expectedTask: &DefaultTask{Title: "renew id card", Status: true},
		},
		{
			name: "fail to test create new task when title is empty",
			actualTask: func() (t Task, e error) {
				return NewTaskFactory().Create("", "", false)
			},
			expectedError: errors.New("title cannot be empty"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task, err := testCase.actualTask()
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

func TestUpdateTaskTitle(t *testing.T) {
	testCases := []struct {
		name          string
		actualTitle   func() string
		expectedTitle string
	}{
		{
			name: "change task title from buy bread to buy groceries",
			actualTitle: func() string {
				task, err := NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)
				task.UpdateTitle("buy groceries")
				return task.GetTitle()
			},
			expectedTitle: "buy groceries",
		},
		{
			name: "change task title from read to read xyz",
			actualTitle: func() string {
				task, err := NewTaskFactory().Create("read", "", false)
				require.NoError(t, err)
				task.UpdateTitle("read xyz")
				return task.GetTitle()
			},
			expectedTitle: "read xyz",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedTitle, testCase.actualTitle())
		})
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	testCases := []struct {
		name                string
		actualDescription   func() string
		expectedDescription string
	}{
		{
			name: "change task description from empty to buy bread, biscuit",
			actualDescription: func() string {
				task, err := NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)
				task.UpdateDescription("buy bread, biscuit")
				return task.GetDescription()
			},
			expectedDescription: "buy bread, biscuit",
		},
		{
			name: "change task description from read book to read xyz",
			actualDescription: func() string {
				task, err := NewTaskFactory().Create("read", "read book", false)
				require.NoError(t, err)
				task.UpdateDescription("read xyz")
				return task.GetDescription()
			},
			expectedDescription: "read xyz",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedDescription, testCase.actualDescription())
		})
	}
}

func TestUpdateTaskTags(t *testing.T) {
	testCases := []struct {
		name         string
		actualTags   func() []string
		expectedTags []string
	}{
		{
			name: "change task tags from empty to groceries, weekly",
			actualTags: func() []string {
				task, err := NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)
				task.UpdateTags("groceries", "weekly")
				return task.GetTags()
			},
			expectedTags: []string{"groceries", "weekly"},
		},
		{
			name: "change task description from weekly, work to daily, work",
			actualTags: func() []string {
				task, err := NewTaskFactory().Create("read", "read book", false, "weekly", "work")
				require.NoError(t, err)
				task.UpdateTags("daily", "work")
				return task.GetTags()
			},
			expectedTags: []string{"daily", "work"},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedTags, testCase.actualTags())
		})
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	testCases := []struct {
		name           string
		actualStatus   func() bool
		expectedStatus bool
	}{
		{
			name: "change task completion Status to completed",
			actualStatus: func() bool {
				task, err := NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)
				task.UpdateStatus()
				return task.GetStatus()
			},
			expectedStatus: true,
		},
		{
			name: "change task completion Status to in-complete",
			actualStatus: func() bool {
				task, err := NewTaskFactory().Create("buy groceries", "", true)
				require.NoError(t, err)
				task.UpdateStatus()
				return task.GetStatus()
			},
			expectedStatus: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedStatus, testCase.actualStatus())
		})
	}
}

func TestGetTaskTitle(t *testing.T) {
	testCases := []struct {
		name          string
		actualTitle   func() string
		expectedTitle string
	}{
		{
			name: "return task Description as buy groceries",
			actualTitle: func() string {
				task, err := NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)
				return task.GetTitle()
			},
			expectedTitle: "buy groceries",
		},
		{
			name: "return task Description as read Xyz",
			actualTitle: func() string {
				task, err := NewTaskFactory().Create("read Xyz", "", false)
				require.NoError(t, err)
				return task.GetTitle()
			},
			expectedTitle: "read Xyz",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedTitle, testCase.actualTitle())
		})
	}
}

func TestGetTaskDescription(t *testing.T) {
	testCases := []struct {
		name                string
		actualDescription   func() string
		expectedDescription string
	}{
		{
			name: "return task Description as buy groceries",
			actualDescription: func() string {
				task, err := NewTaskFactory().Create("buy groceries", "bread, biscuit", false)
				require.NoError(t, err)
				return task.GetDescription()
			},
			expectedDescription: "bread, biscuit",
		},
		{
			name: "return task Description as read Xyz",
			actualDescription: func() string {
				task, err := NewTaskFactory().Create("read", "book, xyz", false)
				require.NoError(t, err)
				return task.GetDescription()
			},
			expectedDescription: "book, xyz",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedDescription, testCase.actualDescription())
		})
	}
}

func TestGetTaskStatus(t *testing.T) {
	testCases := []struct {
		name           string
		actualStatus   func() bool
		expectedStatus bool
	}{
		{
			name: "return Status as false",
			actualStatus: func() bool {
				task, err := NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)
				return task.GetStatus()
			},
			expectedStatus: false,
		},
		{
			name: "return Status as true",
			actualStatus: func() bool {
				task, err := NewTaskFactory().Create("read Xyz", "", true)
				require.NoError(t, err)
				return task.GetStatus()
			},
			expectedStatus: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedStatus, testCase.actualStatus())
		})
	}
}

func TestGetTaskID(t *testing.T) {
	testCases := []struct {
		name     string
		actualID func() string
	}{
		{
			name: "return Status as false",
			actualID: func() string {
				task, err := NewTaskFactory().Create("buy groceries", "", false)
				require.NoError(t, err)
				return task.GetID()
			},
		},
		{
			name: "return Status as true",
			actualID: func() string {
				task, err := NewTaskFactory().Create("read Xyz", "", true)
				require.NoError(t, err)
				return task.GetID()
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, len(testCase.actualID()) == 36)
		})
	}
}

func TestGetTags(t *testing.T) {
	testCases := []struct {
		name         string
		actualTags   func() []string
		expectedTags []string
	}{
		{
			name: "test return empty tags",
			actualTags: func() []string {
				task, err := NewTaskFactory().Create("buy bread", "", false)
				require.NoError(t, err)
				return task.GetTags()
			},
			expectedTags: []string(nil),
		},
		{
			name: "test return daily, work tags",
			actualTags: func() []string {
				task, err := NewTaskFactory().Create("read", "", false, "daily", "work")
				require.NoError(t, err)
				return task.GetTags()
			},
			expectedTags: []string{"daily", "work"},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedTags, testCase.actualTags())
		})
	}
}
