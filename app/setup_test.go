package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"todolist/config"
)

func TestSetUpDependencies(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name             string
		actualDependency Dependencies
	}{
		{
			name:             "create new dependency",
			actualDependency: SetUpDependencies(),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.NotEqual(t, nil, testCase.actualDependency)
			assert.NotEqual(t, nil, testCase.actualDependency.GetTaskService())
		})
	}
}
