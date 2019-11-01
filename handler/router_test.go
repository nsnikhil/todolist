package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/app"
	"todolist/constants"
	"todolist/contract"
	"todolist/domain"
)

func TestRouter(t *testing.T) {
	dependency := app.SetUpDependencies()
	server := httptest.NewServer(NewRouter(dependency))

	testCases := []struct {
		name               string
		actualResponseCode func() int
	}{
		{
			name: "test ping route",
			actualResponseCode: func() int {
				server := httptest.NewServer(NewRouter(app.Dependencies{}))
				response, err := http.Get(fmt.Sprintf("%s/ping", server.URL))
				require.NoError(t, err)
				return response.StatusCode
			},
		},
		{
			name: "test add task route",
			actualResponseCode: func() int {
				b, _ := json.Marshal(domain.MockTask{})
				response, err := http.Post(fmt.Sprintf("%s/todolist/add", server.URL), constants.ContentTypeJSON, bytes.NewBuffer(b))
				require.NoError(t, err)
				return response.StatusCode
			},
		},
		{
			name: "test remove task route",
			actualResponseCode: func() int {
				b, _ := json.Marshal(contract.NewTaskIDRequest(""))
				client := http.Client{}
				request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/todolist/remove", server.URL), bytes.NewReader(b))
				response, err := client.Do(request)
				require.NoError(t, err)
				return response.StatusCode
			},
		},
		{
			name: "test update task route",
			actualResponseCode: func() int {
				b, _ := json.Marshal(domain.MockTask{})
				response, err := http.Post(fmt.Sprintf("%s/todolist/update", server.URL), constants.ContentTypeJSON, bytes.NewBuffer(b))
				require.NoError(t, err)
				return response.StatusCode
			},
		},
		{
			name: "test get todolist route",
			actualResponseCode: func() int {
				response, err := http.Get(fmt.Sprintf("%s/todolist/get-all", server.URL))
				require.NoError(t, err)
				return response.StatusCode
			},
		},
		{
			name: "test get task route",
			actualResponseCode: func() int {
				response, err := http.Get(fmt.Sprintf("%s/todolist/get-one", server.URL))
				require.NoError(t, err)
				return response.StatusCode
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.NotEqual(t, http.StatusNotFound, testCase.actualResponseCode())
		})
	}
}
