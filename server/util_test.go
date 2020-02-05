package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64ToByteSlice(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []byte
		expectedResult []byte
	}{
		{
			name: "test convert 1 to byte slice",
			actualResult: func() []byte {
				return toByteSlice(int64(1))
			},
			expectedResult: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		{
			name: "test convert 8 to byte slice",
			actualResult: func() []byte {
				return toByteSlice(int64(8))
			},
			expectedResult: []byte{0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
