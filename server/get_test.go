package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"todolist/app"
	"todolist/domain"
	"todolist/proto"
	"todolist/service"
)

func TestGetAPI(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)

	testCases := []struct {
		name             string
		actualResponse   func() (*proto.ApiResponse, error)
		expectedResponse func() *proto.ApiResponse
		expectedError    error
	}{
		{
			name: "test get tasks",
			actualResponse: func() (*proto.ApiResponse, error) {
				tasks := []domain.Task{&domain.MockTask{}}

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("GetTasks", []string{"uuid"}).Return(tasks, nil)

				dependency := app.Dependencies{
					Service:     service.NewService(mockTaskService),
					TaskFactory: &domain.MockTaskFactory{},
				}

				server := NewServer(dependency)

				request := &proto.GetRequest{Id: []string{"uuid"}}

				return server.Get(ctx, request)
			},
			expectedResponse: func() *proto.ApiResponse {
				tasks := []domain.Task{&domain.MockTask{}}
				b, err := json.Marshal(tasks)
				require.NoError(t, err)
				return &proto.ApiResponse{Data: &any.Any{Value: b}}
			},
		},
		{
			name: "test get tasks failed",
			actualResponse: func() (*proto.ApiResponse, error) {
				tasks := []domain.Task{&domain.MockTask{}}

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("GetTasks", []string{"uuid"}).Return(tasks, errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(mockTaskService),
					TaskFactory: &domain.MockTaskFactory{},
				}

				server := NewServer(dependency)

				request := &proto.GetRequest{Id: []string{"uuid"}}

				return server.Get(ctx, request)
			},
			expectedError: errors.New("some error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := testCase.actualResponse()
			assert.Equal(t, testCase.expectedError, err)
			if err == nil {
				assert.Equal(t, testCase.expectedResponse(), resp)
			}
		})
	}
}
