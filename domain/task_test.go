package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/apperror"
)

func TestCreateNewTask(t *testing.T) {
	idOne := uuid.New().String()
	idTwo := uuid.New().String()

	testCases := []struct {
		name          string
		actualTask    func() (*Task, error)
		expectedTask  *Task
		expectedError error
	}{
		{
			name: "test create new task",
			actualTask: func() (t *Task, e error) {
				return NewTask(idOne, "buy groceries", false)
			},
			expectedTask: &Task{Id: idOne, Description: "buy groceries", Status: false},
		},
		{
			name: "test create completed task",
			actualTask: func() (t *Task, e error) {
				return NewTask(idTwo, "renew Id card", true)
			},
			expectedTask: &Task{Id: idTwo, Description: "renew Id card", Status: true},
		},
		{
			name: "fail to test create new task when Description is empty",
			actualTask: func() (t *Task, e error) {
				return NewTask(uuid.New().String(), "", false)
			},
			expectedError: apperror.InvalidTaskError{ErrorFormat: "%s : task cannot have empty description", Arguments: []string{"[Task] [NewTask]"}},
		},
		{
			name: "fail to test create new task when uuid is invalid",
			actualTask: func() (t *Task, e error) {
				return NewTask("invalid-id", "call xyz", false)
			},
			expectedError: apperror.InvalidTaskError{ErrorFormat: "%s : invalid uuid %s for the task", Arguments: []string{"[Task] [NewTask]", "invalid-id"}},
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

func TestChangeTaskDescription(t *testing.T) {
	testCases := []struct {
		name         string
		actualTask   func() *Task
		expectedTask *Task
	}{
		{
			name: "change task Description from buy bread to buy groceries",
			actualTask: func() *Task {
				task, _ := NewTask(uuid.New().String(), "buy bread", false)
				task.ChangeDescription("buy groceries")
				return task
			},
			expectedTask: &Task{Description: "buy groceries", Status: false},
		},
		{
			name: "change task Description from read to read xyz",
			actualTask: func() *Task {
				task, _ := NewTask(uuid.New().String(), "read", false)
				task.ChangeDescription("read xyz")
				return task
			},
			expectedTask: &Task{Description: "read xyz", Status: false},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task := testCase.actualTask()
			assert.Equal(t, testCase.expectedTask.Description, task.Description)
			assert.Equal(t, testCase.expectedTask.Status, task.Status)
		})
	}
}

func TestChangeTaskStatus(t *testing.T) {
	testCases := []struct {
		name         string
		actualTask   func() *Task
		expectedTask *Task
	}{
		{
			name: "change task completion Status to completed",
			actualTask: func() *Task {
				task, _ := NewTask(uuid.New().String(), "buy bread", false)
				task.ChangeStatus()
				return task
			},
			expectedTask: &Task{Description: "buy bread", Status: true},
		},
		{
			name: "change task completion Status to in-complete",
			actualTask: func() *Task {
				task, _ := NewTask(uuid.New().String(), "buy groceries", true)
				task.ChangeStatus()
				return task
			},
			expectedTask: &Task{Description: "buy groceries", Status: false},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			task := testCase.actualTask()
			assert.Equal(t, testCase.expectedTask.Description, task.Description)
			assert.Equal(t, testCase.expectedTask.Status, task.Status)
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
				task, _ := NewTask(uuid.New().String(), "buy groceries", false)
				return task.GetDescription()
			},
			expectedDescription: "buy groceries",
		},
		{
			name: "return task Description as read Xyz",
			actualDescription: func() string {
				task, _ := NewTask(uuid.New().String(), "read Xyz", false)
				return task.GetDescription()
			},
			expectedDescription: "read Xyz",
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
				task, _ := NewTask(uuid.New().String(), "buy groceries", false)
				return task.GetStatus()
			},
			expectedStatus: false,
		},
		{
			name: "return Status as true",
			actualStatus: func() bool {
				task, _ := NewTask(uuid.New().String(), "read Xyz", true)
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
	idOne := uuid.New().String()
	idTwo := uuid.New().String()

	testCases := []struct {
		name       string
		actualID   func() string
		expectedID string
	}{
		{
			name: "return Status as false",
			actualID: func() string {
				task, _ := NewTask(idOne, "buy groceries", false)
				return task.GetID()
			},
			expectedID: idOne,
		},
		{
			name: "return Status as true",
			actualID: func() string {
				task, _ := NewTask(idTwo, "read Xyz", true)
				return task.GetID()
			},
			expectedID: idTwo,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedID, testCase.actualID())
		})
	}
}

func TestValidateTask(t *testing.T) {
	testCases := []struct {
		name           string
		actualOutput   func() bool
		expectedOutput bool
	}{
		{
			name: "return true when all details are correct",
			actualOutput: func() bool {
				var task Task
				task.Id = uuid.New().String()
				task.Description = "buy biscuits"
				return task.Validate()
			},
			expectedOutput: true,
		},
		{
			name: "return true when all description is empty",
			actualOutput: func() bool {
				var task Task
				task.Id = uuid.New().String()
				return task.Validate()
			},
		},
		{
			name: "return true when all id is invalid uuid",
			actualOutput: func() bool {
				var task Task
				task.Description = "read abc"
				task.Id = "invalid-uuid"
				return task.Validate()
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedOutput, testCase.actualOutput())
		})
	}
}
