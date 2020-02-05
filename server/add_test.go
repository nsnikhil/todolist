package server

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todolist/app"
	"todolist/domain"
	"todolist/proto"
	"todolist/service"
)

func TestAddAPI(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)

	testCases := []struct {
		name             string
		actualResponse   func() (*proto.ApiResponse, error)
		expectedResponse *proto.ApiResponse
		expectedError    error
	}{
		{
			name: "test success response",
			actualResponse: func() (*proto.ApiResponse, error) {

				mockTask := &domain.MockTask{}
				mockTask.On("GetTitle").Return("title")
				mockTask.On("GetStatus").Return(false)

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Add", mockTask).Return("some-id", nil)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("Create", "title", "", false, []string(nil)).Return(mockTask, nil)

				appService := service.NewService(mockTaskService)

				dependency := app.Dependencies{
					Service:     appService,
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.AddRequest{Title: mockTask.GetTitle(), Status: mockTask.GetStatus()}

				return server.Add(ctx, request)
			},
			expectedResponse: &proto.ApiResponse{
				Data: &any.Any{
					TypeUrl: "string",
					Value:   []byte("some-id"),
				},
			},
		},
		{
			name: "test new task creation failed",
			actualResponse: func() (*proto.ApiResponse, error) {

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("Create", "", "", false, []string(nil)).Return(&domain.MockTask{}, errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(&service.MockTaskService{}),
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.AddRequest{Title: "", Status: false}

				return server.Add(ctx, request)
			},
			expectedResponse: &proto.ApiResponse{},
			expectedError:    errors.New("some error"),
		},
		{
			name: "test task service failed to add task",
			actualResponse: func() (*proto.ApiResponse, error) {
				mockTask := &domain.MockTask{}
				mockTask.On("GetTitle").Return("title")
				mockTask.On("GetStatus").Return(false)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("Create", "", "", false, []string(nil)).Return(mockTask, nil)

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Add", mockTask).Return("some-id", errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(mockTaskService),
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.AddRequest{Title: "", Status: false}

				return server.Add(ctx, request)
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
