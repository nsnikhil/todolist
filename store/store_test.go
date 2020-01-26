package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewStore(t *testing.T) {
	testCases := []struct {
		name          string
		actualStore   func() Store
		expectedStore Store
	}{
		{
			name: "test create new store",
			actualStore: func() Store {
				return NewStore(&sqlx.DB{})
			},
			expectedStore: Store{TaskStore: NewTaskStore(&sqlx.DB{})},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedStore, testCase.actualStore())
		})
	}
}
