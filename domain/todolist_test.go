package domain

import (
	uuid2 "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/apperror"
)

func TestCreateNewTodoList(t *testing.T) {
	testCases := []struct {
		name             string
		actualTodoList   func() *TodoList
		expectedTodoList *TodoList
		expectedError    error
	}{
		{
			name: "test create new todo list with a new task",
			actualTodoList: func() *TodoList {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				return NewTodoList(buyGroceries)
			},
			expectedTodoList: &TodoList{Tasks: []*Task{{Description: "buy groceries", Status: false}}},
		},
		{
			name: "test create new todo list with a two new task",
			actualTodoList: func() *TodoList {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				readXyz, _ := NewTask(uuid2.New().String(), "read Xyz", false)
				return NewTodoList(buyGroceries, readXyz)
			},
			expectedTodoList: &TodoList{Tasks: []*Task{
				{Description: "read Xyz", Status: false},
				{Description: "buy groceries", Status: false},
			}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			todoList := testCase.actualTodoList()
			for i := 0; i < len(testCase.expectedTodoList.Tasks); i++ {
				assert.Equal(t, testCase.expectedTodoList.Tasks[i].Description, todoList.Tasks[i].Description)
			}
		})
	}
}

func TestAddTaskToTodoList(t *testing.T) {
	testCases := []struct {
		name             string
		actualTodoList   func() *TodoList
		expectedTodoList *TodoList
	}{
		{
			name: "add one task to todo list",
			actualTodoList: func() *TodoList {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				list := NewTodoList(buyGroceries)
				readXyz, _ := NewTask(uuid2.New().String(), "read Xyz", false)
				list.Add(readXyz)
				return list
			},
			expectedTodoList: &TodoList{Tasks: []*Task{
				{Description: "buy groceries", Status: false},
				{Description: "read Xyz", Status: false},
			}},
		},
		{
			name: "add two plus one task to todo list",
			actualTodoList: func() *TodoList {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				list := NewTodoList(buyGroceries)
				readXyz, _ := NewTask(uuid2.New().String(), "read Xyz", false)
				list.Add(readXyz)
				cookLunch, _ := NewTask(uuid2.New().String(), "cook lunch", false)
				list.Add(cookLunch)
				callPerson, _ := NewTask(uuid2.New().String(), "call person", false)
				list.Add(callPerson)
				return list
			},
			expectedTodoList: &TodoList{Tasks: []*Task{
				{Description: "buy groceries", Status: false},
				{Description: "read Xyz", Status: false},
				{Description: "cook lunch", Status: false},
				{Description: "call person", Status: false},
			}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			todoList := testCase.actualTodoList()
			for i := 0; i < len(testCase.expectedTodoList.Tasks); i++ {
				assert.Equal(t, testCase.expectedTodoList.Tasks[i].Description, todoList.Tasks[i].Description)
			}
		})
	}
}

func TestRemoveTaskFromTodoList(t *testing.T) {
	testCases := []struct {
		name             string
		actualTodoList   func() (*TodoList, error)
		expectedTodoList *TodoList
		expectedError    error
	}{
		{
			name: "test remove task from todo list with one task",
			actualTodoList: func() (*TodoList, error) {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				list := NewTodoList(buyGroceries)
				err := list.Remove(buyGroceries)
				return list, err
			},
			expectedTodoList: &TodoList{Tasks: []*Task{}},
		},
		{
			name: "test remove task from todo list with two task",
			actualTodoList: func() (*TodoList, error) {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				readXyz, _ := NewTask(uuid2.New().String(), "read Xyz", false)
				list := NewTodoList(buyGroceries, readXyz)
				err := list.Remove(buyGroceries)
				return list, err
			},
			expectedTodoList: &TodoList{Tasks: []*Task{{Description: "read Xyz", Status: false}}},
		},
		{
			name: "test fail to remove non present task from todo list",
			actualTodoList: func() (*TodoList, error) {
				buyGroceries, _ := NewTask(uuid2.New().String(), "buy groceries", false)
				readXyz, _ := NewTask(uuid2.New().String(), "read Xyz", false)
				list := NewTodoList(buyGroceries)
				err := list.Remove(readXyz)
				return list, err
			},
			expectedTodoList: &TodoList{Tasks: []*Task{{Description: "buy groceries", Status: false}}},
			expectedError:    apperror.TaskRemoveFailedError{ErrorFormat: "read Xyz"},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			todoList, err := testCase.actualTodoList()
			assert.Equal(t, testCase.expectedError, err)
			if err != nil {
				for i := 0; i < len(testCase.expectedTodoList.Tasks); i++ {
					assert.Equal(t, testCase.expectedTodoList.Tasks[i].Description, todoList.Tasks[i].Description)
				}
			}
		})
	}
}
