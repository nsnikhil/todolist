package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/health/grpc_health_v1"
	"testing"
	"todolist/proto"
)

//noinspection GoErrorStringFormat
func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (*grpc_health_v1.HealthCheckResponse, error)
		expectedResult *grpc_health_v1.HealthCheckResponse
		expectedError  error
	}{
		{
			name: "test return serving",
			actualResult: func() (*grpc_health_v1.HealthCheckResponse, error) {
				mockResp := &proto.ApiResponse{
					Data: &any.Any{TypeUrl: "string",
						Value: []byte("pong"),
					},
				}

				mockServer := &MockAPIServer{}
				mockServer.On("Ping", context.Background(), &proto.PingRequest{}).Return(mockResp, nil)

				healthServer := NewHealthServer(mockServer)

				return healthServer.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
			},
			expectedResult: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_SERVING,
			},
		},
		{
			name: "test return not serving when ping fails",
			actualResult: func() (*grpc_health_v1.HealthCheckResponse, error) {
				mockServer := &MockAPIServer{}
				mockServer.On("Ping", context.Background(), &proto.PingRequest{}).Return(&proto.ApiResponse{}, errors.New("some error"))

				healthServer := NewHealthServer(mockServer)

				return healthServer.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
			},
			expectedResult: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
			},
			expectedError: errors.New("some error"),
		},
		{
			name: "test return not serving when ping return nil data",
			actualResult: func() (*grpc_health_v1.HealthCheckResponse, error) {
				mockServer := &MockAPIServer{}
				mockServer.On("Ping", context.Background(), &proto.PingRequest{}).Return(&proto.ApiResponse{}, nil)

				healthServer := NewHealthServer(mockServer)

				return healthServer.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
			},
			expectedResult: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
			},
			expectedError: fmt.Errorf("invalid response :"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
