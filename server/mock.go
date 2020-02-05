package server

import (
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/health/grpc_health_v1"
	"todolist/proto"
)

type MockAPIServer struct {
	mock.Mock
}

func (m *MockAPIServer) Ping(ctx context.Context, in *proto.PingRequest) (*proto.ApiResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.ApiResponse), args.Error(1)
}

func (m *MockAPIServer) Add(ctx context.Context, in *proto.AddRequest) (*proto.ApiResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.ApiResponse), args.Error(1)
}

func (m *MockAPIServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.ApiResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.ApiResponse), args.Error(1)
}

func (m *MockAPIServer) Remove(ctx context.Context, in *proto.RemoveRequest) (*proto.ApiResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.ApiResponse), args.Error(1)
}

func (m *MockAPIServer) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.ApiResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*proto.ApiResponse), args.Error(1)
}

type MockHealthServer struct {
	mock.Mock
}

func (m MockHealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*grpc_health_v1.HealthCheckResponse), args.Error(1)
}

func (m MockHealthServer) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	args := m.Called(req, w)
	return args.Error(0)
}
