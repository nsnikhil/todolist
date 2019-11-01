package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/contract"
	"todolist/service"
)

func TestRemoveHandlerSuccess(t *testing.T) {
	id := uuid.New().String()
	removeTaskRequest := contract.NewTaskIDRequest(id)

	b, err := json.Marshal(removeTaskRequest)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/todolist/remove", bytes.NewReader(b))

	mockTodoListService := service.MockTodoListServiceInterface{}
	mockTodoListService.On("Remove", id).Return(nil)

	removeHandler(mockTodoListService)(w, r)

	expectedResponseBody := "{\"success\":true,\"error\":\"\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestRemoveHandlerErrorCases(t *testing.T) {
	testCases := []struct {
		name                    string
		actualResponse          func() (int, string)
		expectedResponseCode    int
		expectedResponseMessage string
	}{
		{
			name: "fail to unmarshal invalid request body",
			actualResponse: func() (int, string) {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/todolist/remove", nil)

				removeHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"unexpected end of JSON input\"}\n",
		},
		{
			name: "fail to validate request body",
			actualResponse: func() (int, string) {
				var invalidTask contract.TaskIDRequest
				invalidTask.TaskID = "invalid uuid"

				b, err := json.Marshal(invalidTask)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/todolist/remove", bytes.NewReader(b))

				removeHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"[removeHandler] [Validate] : failed to validate task invalid uuid\"}\n",
		},
		{
			name: "fail to remove task",
			actualResponse: func() (int, string) {
				id := uuid.New().String()
				removeTaskRequest := contract.NewTaskIDRequest(id)

				b, err := json.Marshal(removeTaskRequest)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/todolist/remove", bytes.NewReader(b))

				mockTodoListService := service.MockTodoListServiceInterface{}
				mockTodoListService.On("Remove", id).Return(errors.New("failed to remove task"))

				removeHandler(mockTodoListService)(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusInternalServerError,
			expectedResponseMessage: "{\"success\":false,\"error\":\"failed to remove task\"}\n",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			statusCode, errorMessage := testCase.actualResponse()
			assert.Equal(t, testCase.expectedResponseCode, statusCode)
			assert.Equal(t, testCase.expectedResponseMessage, errorMessage)
		})
	}
}
