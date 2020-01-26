package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/app"
)

func TestNewServer(t *testing.T) {
	actualServer := NewServer(app.Dependencies{})
	expectedServer := Server{app.Dependencies{}}
	assert.Equal(t, expectedServer, actualServer)
}
