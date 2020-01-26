package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"todolist/app"
	"todolist/domain"
	"todolist/proto"
	"todolist/service"
)

func TestPing(t *testing.T) {
	dependency := app.Dependencies{
		Service: service.NewService(&service.MockTaskService{}),
		TaskFactory: &domain.MockTaskFactory{},
	}
	server := NewServer(dependency)
	response, err := server.Ping(context.Background(), &proto.PingRequest{})
	require.NoError(t, err)
	assert.Equal(t, "pong", string(response.Data.Value))
}
