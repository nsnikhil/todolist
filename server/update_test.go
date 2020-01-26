package server

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todolist/app"
	"todolist/domain"
	"todolist/proto"
	"todolist/service"
)

func TestUpdateAPI(t *testing.T) {
	id := uuid.New().String()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)

	testCases := []struct {
		name             string
		actualResponse   func() (*proto.ApiResponse, error)
		expectedResponse *proto.ApiResponse
		expectedError    error
	}{
		{
			name: "test update task",
			actualResponse: func() (*proto.ApiResponse, error) {
				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return(id)
				mockTask.On("GetTitle").Return("title")
				mockTask.On("GetStatus").Return(false)

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Update", mockTask).Return(nil)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("CreateWithID", id, "title", "", false, []string(nil)).Return(mockTask, nil)

				appService := service.NewService(mockTaskService)

				dependency := app.Dependencies{
					Service:     appService,
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.UpdateRequest{
					Id:     mockTask.GetID(),
					Title:  mockTask.GetTitle(),
					Status: mockTask.GetStatus(),
				}

				return server.Update(ctx, request)
			},
			expectedResponse: &proto.ApiResponse{},
		},
		{
			name: "test update task failed because of invalid uuid",
			actualResponse: func() (*proto.ApiResponse, error) {
				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return("invalid id")
				mockTask.On("GetTitle").Return("title")
				mockTask.On("GetStatus").Return(false)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("CreateWithID", "invalid id", "title", "", false, []string(nil)).Return(mockTask, errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(&service.MockTaskService{}),
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.UpdateRequest{
					Id:     mockTask.GetID(),
					Title:  mockTask.GetTitle(),
					Status: mockTask.GetStatus(),
				}

				return server.Update(ctx, request)
			},
			expectedError:errors.New("some error"),
		},
		{
			name: "test update task failed because of empty title",
			actualResponse: func() (*proto.ApiResponse, error) {
				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return(id)
				mockTask.On("GetTitle").Return("")
				mockTask.On("GetStatus").Return(false)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("CreateWithID", id, "", "", false, []string(nil)).Return(mockTask, errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(&service.MockTaskService{}),
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.UpdateRequest{
					Id:     mockTask.GetID(),
					Title:  mockTask.GetTitle(),
					Status: mockTask.GetStatus(),
				}

				return server.Update(ctx, request)
			},
			expectedError:errors.New("some error"),
		},
		{
			name: "test task service failed to update task",
			actualResponse: func() (*proto.ApiResponse, error) {
				mockTask := &domain.MockTask{}
				mockTask.On("GetID").Return(id)
				mockTask.On("GetTitle").Return("title")
				mockTask.On("GetStatus").Return(false)

				mockFactory := &domain.MockTaskFactory{}
				mockFactory.On("CreateWithID", id, "title", "", false, []string(nil)).Return(mockTask, nil)

				mockTaskService := &service.MockTaskService{}
				mockTaskService.On("Update", mockTask).Return(errors.New("some error"))

				dependency := app.Dependencies{
					Service:     service.NewService(mockTaskService),
					TaskFactory: mockFactory,
				}

				server := NewServer(dependency)

				request := &proto.UpdateRequest{
					Id:     mockTask.GetID(),
					Title:  mockTask.GetTitle(),
					Status: mockTask.GetStatus(),
				}

				return server.Update(ctx, request)
			},
			expectedError:errors.New("some error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := testCase.actualResponse()
			assert.Equal(t, testCase.expectedError, err)
			if err == nil {
				assert.Equal(t, testCase.expectedResponse, resp)
			}
		})
	}
}
