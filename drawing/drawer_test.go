package drawing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPath(t *testing.T) {
	testCases := []struct {
		input     uint64
		expOutput string
	}{{
		input:     0,
		expOutput: `M 0 0 `,
	}, {
		input:     1,
		expOutput: `M 0 0 l 1 0 `,
	}, {
		input:     2,
		expOutput: `M 0 0 l 1 0 l 0 -1 `,
	}, {
		input:     3,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 `,
	}, {
		input:     4,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 l 0 -1 `,
	}, {
		input:     5,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 l 0 -1 l -1 0 `,
	}, {
		input:     6,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 l 0 -1 l -1 0 l 0 1 `,
	}, {
		input:     7,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 l 0 -1 l -1 0 l 0 1 l -1 0 `,
	}, {
		input:     8,
		expOutput: `M 0 0 l 1 0 l 0 -1 l -1 0 l 0 -1 l -1 0 l 0 1 l -1 0 l 0 -1 `,
	}}

	for _, tc := range testCases {
		actOutput := buildPath(tc.input)
		assert.Equal(t, tc.expOutput, actOutput, `received unexpected output for "%d"`, tc.input)
	}
}
