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

func TestAddHandlerSuccess(t *testing.T) {
	task, _ := domain.NewTask(uuid.New().String(), "buy biscuits", false)

	b, err := json.Marshal(task)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/todolist/add", bytes.NewReader(b))

	mockTodoListService := service.MockTodoListServiceInterface{}
	mockTodoListService.On("Add", task).Return(nil)

	addHandler(mockTodoListService)(w, r)

	expectedResponseBody := "{\"success\":true,\"error\":\"\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestAddHandlerErrorCases(t *testing.T) {
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
				r := httptest.NewRequest(http.MethodPost, "/todolist/add", nil)

				addHandler(service.MockTodoListServiceInterface{})(w, r)

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
				r := httptest.NewRequest(http.MethodPost, "/todolist/add", bytes.NewReader(b))

				addHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"[addHandler] [Validate] : failed to validate task {invalid uuid buy bread false}\"}\n",
		},
		{
			name: "fail to add task",
			actualResponse: func() (int, string) {
				task, _ := domain.NewTask(uuid.New().String(), "buy biscuits", false)

				b, err := json.Marshal(task)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/todolist/add", bytes.NewReader(b))

				mockTodoListService := service.MockTodoListServiceInterface{}
				mockTodoListService.On("Add", task).Return(errors.New("failed to add task"))

				addHandler(mockTodoListService)(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusInternalServerError,
			expectedResponseMessage: "{\"success\":false,\"error\":\"failed to add task\"}\n",
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
