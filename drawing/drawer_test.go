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
		expOutput: `M ` + startX + ` ` + startY + ` `,
	}, {
		input:     1,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` `,
	}, {
		input:     2,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` `,
	}, {
		input:     3,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` `,
	}, {
		input:     4,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` v -` + delta + ` `,
	}, {
		input:     5,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` v -` + delta + ` h -` + delta + ` `,
	}, {
		input:     6,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` v -` + delta + ` h -` + delta + ` v ` + delta + ` `,
	}, {
		input:     7,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` v -` + delta + ` h -` + delta + ` v ` + delta + ` h -` + delta + ` `,
	}, {
		input:     8,
		expOutput: `M ` + startX + ` ` + startY + ` h ` + delta + ` v -` + delta + ` h -` + delta + ` v -` + delta + ` h -` + delta + ` v ` + delta + ` h -` + delta + ` v -` + delta + ` `,
	}}

	for _, tc := range testCases {
		actOutput := BuildPath(tc.input)
		assert.Equal(t, tc.expOutput, actOutput, `received unexpected output for "%d"`, tc.input)
	}
}
