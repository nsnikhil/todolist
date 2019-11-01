package service

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreateNewService(t *testing.T) {
	testCases := []struct {
		name            string
		actualService   Service
		expectedService Service
	}{
		{
			name:            "test create new service",
			actualService:   NewService(MockTodoListServiceInterface{}),
			expectedService: Service{todoListService: MockTodoListServiceInterface{}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedService, testCase.actualService)
		})
	}
}
