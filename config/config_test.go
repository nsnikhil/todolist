package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewConfig(t *testing.T) {
	err := Load()
	require.NoError(t, err)

	testCases := []struct {
		name           string
		actualConfig   config
		expectedConfig config
	}{
		{
			name:         "test create new app config",
			actualConfig: appConfig,
			expectedConfig: config{
				databaseConfig: newDatabaseConfig(),
				serverConfig:   newSeverConfig(),
				logConfig:      newLogConfig(),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedConfig, testCase.actualConfig)
		})
	}
}
