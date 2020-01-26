package server

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todolist/app"
	"todolist/domain"
	"todolist/proto"
	"todolist/service"
)

func TestRemoveAPI(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)

	testCases := []struct {
		name             string
		actualResponse   func() (*proto.ApiResponse, error)
		expectedResponse *proto.ApiResponse
		expectedError    error
	}{
		{
			name: "test remove one task",
			actualResponse: func() (*proto.ApiResponse, error) {
				id := "uuid"

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Remove", id, []string(nil)).Return(nil)

				appService := service.NewService(mockTaskService)

				dependency := app.Dependencies{
					Service:     appService,
					TaskFactory: &domain.MockTaskFactory{},
				}

				server := NewServer(dependency)

				request := &proto.RemoveRequest{Id: id}

				return server.Remove(ctx, request)

			},
			expectedResponse: &proto.ApiResponse{},
		},
		{
			name: "test remove multiple task",
			actualResponse: func() (*proto.ApiResponse, error) {
				id := "uuid"

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Remove", id, []string{"other-uuid"}).Return(nil)

				appService := service.NewService(mockTaskService)

				dependency := app.Dependencies{
					Service:     appService,
					TaskFactory: &domain.MockTaskFactory{},
				}

				server := NewServer(dependency)

				request := &proto.RemoveRequest{Id: id, Ids: []string{"other-uuid"}}

				return server.Remove(ctx, request)

			},
			expectedResponse: &proto.ApiResponse{},
		},
		{
			name: "test remove task failed",
			actualResponse: func() (*proto.ApiResponse, error) {
				id := "uuid"

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Remove", id, []string(nil)).Return(errors.New("some error"))

				appService := service.NewService(mockTaskService)

				dependency := app.Dependencies{
					Service:     appService,
					TaskFactory: &domain.MockTaskFactory{},
				}

				server := NewServer(dependency)

				request := &proto.RemoveRequest{Id: id}

				return server.Remove(ctx, request)

			},
			expectedResponse: &proto.ApiResponse{},
			expectedError:    errors.New("some error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := testCase.actualResponse()
			assert.Equal(t, testCase.expectedResponse, resp)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
