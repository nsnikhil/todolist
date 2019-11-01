package contract

import (
	"github.com/google/uuid"
	"gotest.tools/assert"
	"testing"
)

func TestCreateNewTodoListResponse(t *testing.T) {
	testCases := []struct {
		name             string
		actualResponse   func() TodoListResponse
		expectedResponse TodoListResponse
	}{
		{
			name: "test create new success response",
			actualResponse: func() TodoListResponse {
				return NewTodoListResponse(nil, "", true)
			},
			expectedResponse: TodoListResponse{
				Data:         nil,
				Success:      true,
				ErrorMessage: "",
			},
		},
		{
			name: "test create new failure response",
			actualResponse: func() TodoListResponse {
				return NewTodoListResponse(nil, "invalid data", false)
			},
			expectedResponse: TodoListResponse{
				Data:         nil,
				Success:      false,
				ErrorMessage: "invalid data",
			},
		},
		{
			name: "test create new success response with data",
			actualResponse: func() TodoListResponse {
				return NewTodoListResponse(&struct{}{}, "", true)
			},
			expectedResponse: TodoListResponse{
				Data:         &struct{}{},
				Success:      true,
				ErrorMessage: "",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResponse, testCase.actualResponse())
		})
	}
}

func TestCreateNewTaskIdRequest(t *testing.T) {
	idOne := uuid.New().String()
	testCases := []struct {
		name            string
		actualRequest   func() TaskIDRequest
		expectedRequest TaskIDRequest
	}{
		{
			name: "test create new task id request",
			actualRequest: func() TaskIDRequest {
				return NewTaskIDRequest(idOne)
			},
			expectedRequest: TaskIDRequest{TaskID: idOne},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedRequest, testCase.actualRequest())
		})
	}
}

func TestTaskIDRequestValidate(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test validate success",
			actualResult: func() bool {
				request := NewTaskIDRequest(uuid.New().String())
				return request.Validate()
			},
			expectedResult: true,
		},
		{
			name: "test validate failure",
			actualResult: func() bool {
				request := NewTaskIDRequest("invalid uuid")
				return request.Validate()
			},
			expectedResult: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
