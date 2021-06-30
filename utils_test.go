package gowatchprog_test

import (
	"testing"

	"github.com/micaiahwallace/gowatchprog"
	"github.com/stretchr/testify/assert"
)

func TestGetCommandLine(t *testing.T) {

	testCases := []struct {
		name           string
		binPath        string
		args           []string
		expectedResult string
	}{
		{
			name:           `GetCommandLine returns the quoted path alone when empty arguments array is passed`,
			binPath:        `c:\test\windows\path`,
			args:           []string{},
			expectedResult: `"c:\test\windows\path"`,
		},
		{
			name:           `GetCommandLine returns the quoted path with appended arguments when arguments are passed`,
			binPath:        `c:\test\windows\path`,
			args:           []string{"1", "2", "3"},
			expectedResult: `"c:\test\windows\path" 1 2 3`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualResult := gowatchprog.GetCommandLine(tc.binPath, tc.args)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
