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
	"todolist/domain"
	"todolist/service"
)

func TestUpdateHandlerSuccess(t *testing.T) {
	task, _ := domain.NewTask(uuid.New().String(), "buy biscuits", false)

	b, err := json.Marshal(task)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/todolist/update", bytes.NewReader(b))

	mockTodoListService := service.MockTodoListServiceInterface{}
	mockTodoListService.On("Update", task).Return(nil)

	updateHandler(mockTodoListService)(w, r)

	expectedResponseBody := "{\"success\":true,\"error\":\"\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestUpdateHandlerErrorCases(t *testing.T) {
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
				r := httptest.NewRequest(http.MethodPost, "/todolist/update", nil)

				updateHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"unexpected end of JSON input\"}\n",
		},
		{
			name: "fail to validate request body",
			actualResponse: func() (int, string) {
				var invalidTask domain.Task
				invalidTask.Id = "invalid uuid"
				invalidTask.Description = "buy bread"

				b, err := json.Marshal(invalidTask)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/todolist/update", bytes.NewReader(b))

				updateHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"[updateHandler] [Validate] : failed to validate task {invalid uuid buy bread false}\"}\n",
		},
		{
			name: "fail to update task",
			actualResponse: func() (int, string) {
				task, _ := domain.NewTask(uuid.New().String(), "buy biscuits", false)

				b, err := json.Marshal(task)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/todolist/update", bytes.NewReader(b))

				mockTodoListService := service.MockTodoListServiceInterface{}
				mockTodoListService.On("Update", task).Return(errors.New("failed to update task"))

				updateHandler(mockTodoListService)(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusInternalServerError,
			expectedResponseMessage: "{\"success\":false,\"error\":\"failed to update task\"}\n",
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
