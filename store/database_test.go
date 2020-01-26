package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"todolist/config"
)

func TestNewDBHandler(t *testing.T) {
	err := config.Load()
	require.NoError(t, err)

	testCases := []struct {
		name          string
		actualDB      func() (*sqlx.DB, error)
		expectedDB    *sqlx.DB
		expectedError error
	}{
		{
			name: "test create new db handle",
			actualDB: func() (*sqlx.DB, error) {
				dbHandle := NewDBHandler(config.GetDatabaseConfig())
				return dbHandle.GetDB()
			},
			expectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.actualDB()
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
