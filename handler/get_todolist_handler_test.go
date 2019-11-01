package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/domain"
	"todolist/service"
)

func TestGetTodoListHandlerSuccess(t *testing.T) {
	idOne := "18115db6-6ca9-46e7-bded-cae3c76f15e2"
	idTwo := "50f3431e-cf8f-496b-b374-dae2347d343c"

	var buyBread domain.Task
	buyBread.Id = idOne
	buyBread.Description = "buy bread"

	var readXyz domain.Task
	readXyz.Id = idTwo
	readXyz.Description = "read Xyz"

	todoList := domain.NewTodoList(&buyBread, &readXyz)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/todolist/get-all", nil)

	mockTodoListService := service.MockTodoListServiceInterface{}
	mockTodoListService.On("GetTodoList").Return(todoList, nil)

	getTodoListHandler(mockTodoListService)(w, r)

	expectedResponseBody := "{\"data\":{\"Tasks\":[{\"Id\":\"50f3431e-cf8f-496b-b374-dae2347d343c\",\"Description\":\"read Xyz\",\"Status\":false},{\"Id\":\"18115db6-6ca9-46e7-bded-cae3c76f15e2\",\"Description\":\"buy bread\",\"Status\":false}]},\"success\":true,\"error\":\"\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestGetTodoListHandlerErrorCase(t *testing.T) {
	testCases := []struct {
		name                    string
		actualResponse          func() (int, string)
		expectedResponseCode    int
		expectedResponseMessage string
	}{
		{
			name: "fail to get todolist",
			actualResponse: func() (int, string) {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/todolist/get-all", nil)

				mockTodoListService := service.MockTodoListServiceInterface{}
				mockTodoListService.On("GetTodoList").Return(&domain.MockTodoList{}, errors.New("failed to get todolist"))

				getTodoListHandler(mockTodoListService)(w, r)

				return w.Code, w.Body.String()
			},
			expectedResponseCode:    http.StatusInternalServerError,
			expectedResponseMessage: "{\"success\":false,\"error\":\"failed to get todolist\"}\n",
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
