package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/contract"
	"todolist/domain"
	"todolist/service"
)

func TestGetTaskHandlerSuccess(t *testing.T) {
	idOne := "18115db6-6ca9-46e7-bded-cae3c76f15e2"

	var buyBread domain.Task
	buyBread.Id = idOne
	buyBread.Description = "buy bread"

	getTaskRequest := contract.NewTaskIDRequest(idOne)

	b, err := json.Marshal(getTaskRequest)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/todolist/get-one", bytes.NewReader(b))

	mockTodoListService := service.MockTodoListServiceInterface{}
	mockTodoListService.On("GetTask", idOne).Return(&buyBread, nil)

	getTaskHandler(mockTodoListService)(w, r)

	expectedResponseBody := "{\"data\":{\"Id\":\"18115db6-6ca9-46e7-bded-cae3c76f15e2\",\"Description\":\"buy bread\",\"Status\":false},\"success\":true,\"error\":\"\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestGetTaskHandlerErrorCase(t *testing.T) {
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
				r := httptest.NewRequest(http.MethodGet, "/todolist/get-one", nil)

				getTaskHandler(service.MockTodoListServiceInterface{})(w, r)

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
				r := httptest.NewRequest(http.MethodGet, "/todolist/get-one", bytes.NewReader(b))

				getTaskHandler(service.MockTodoListServiceInterface{})(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusBadRequest,
			expectedResponseMessage: "{\"success\":false,\"error\":\"[getTaskHandler] [Validate] : failed to validate task invalid uuid\"}\n",
		},
		{
			name: "fail to get task",
			actualResponse: func() (int, string) {
				idOne := "18115db6-6ca9-46e7-bded-cae3c76f15e2"

				getTaskRequest := contract.NewTaskIDRequest(idOne)

				b, err := json.Marshal(getTaskRequest)
				require.NoError(t, err)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/todolist/get-one", bytes.NewReader(b))

				mockTodoListService := service.MockTodoListServiceInterface{}
				mockTodoListService.On("GetTask", idOne).Return(&domain.MockTask{}, errors.New("failed to get task"))

				getTaskHandler(mockTodoListService)(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusInternalServerError,
			expectedResponseMessage: "{\"success\":false,\"error\":\"failed to get task\"}\n",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			responseCode, responseMessage := testCase.actualResponse()
			assert.Equal(t, testCase.expectedResponseCode, responseCode)
			assert.Equal(t, testCase.expectedResponseMessage, responseMessage)
		})
	}
}
