package gowatchprog_test

import (
	"testing"

	"github.com/micaiahwallace/gowatchprog"
	"github.com/stretchr/testify/assert"
)

func TestSafeName(t *testing.T) {

	testCases := []struct {
		name           string
		inputName      string
		expectedResult string
	}{
		{
			name:           "safeName replaces non-ascii symbols with dashes when they are present",
			inputName:      "Test name!@here123",
			expectedResult: "Test-name--here123",
		},
		{
			name:           "safeName returns the input value when it is valid",
			inputName:      "ok-name-here",
			expectedResult: "ok-name-here",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			p := &gowatchprog.Program{Name: tc.inputName}
			actual := p.SafeName()
			assert.Equal(t, tc.expectedResult, actual)
		})
	}
}
