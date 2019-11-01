package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateNewServer(t *testing.T) {
	testCases := []struct {
		name           string
		actualServer   *Server
		expectedServer *Server
	}{
		{
			name:         "create new server",
			actualServer: NewServer(nil),
			expectedServer: &Server{
				apiServer: &http.Server{
					Handler: nil,
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedServer, testCase.actualServer)
		})
	}
}
