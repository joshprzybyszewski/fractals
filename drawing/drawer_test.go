package drawing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPath(t *testing.T) {
	delta := 1
	deltaStr := `1`

	pb := New(delta)

	testCases := []struct {
		input     uint64
		expOutput string
		expMaxX   int64
		expMaxY   int64
	}{{
		input:     0,
		expOutput: `M 0 0 `,
		expMaxX:   0,
		expMaxY:   0,
	}, {
		input:     1,
		expOutput: `M 0 0 h ` + deltaStr + ` `,
		expMaxX:   1,
		expMaxY:   0,
	}, {
		input:     2,
		expOutput: `M 0 1 h ` + deltaStr + ` v -` + deltaStr + ` `,
		expMaxX:   1,
		expMaxY:   1,
	}, {
		input:     3,
		expOutput: `M 0 1 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` `,
		expMaxX:   1,
		expMaxY:   1,
	}, {
		input:     4,
		expOutput: `M 0 2 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` `,
		expMaxX:   1,
		expMaxY:   2,
	}, {
		input:     5,
		expOutput: `M 1 2 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` `,
		expMaxX:   2,
		expMaxY:   2,
	}, {
		input:     6,
		expOutput: `M 1 2 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v ` + deltaStr + ` `,
		expMaxX:   2,
		expMaxY:   2,
	}, {
		input:     7,
		expOutput: `M 2 2 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v ` + deltaStr + ` h -` + deltaStr + ` `,
		expMaxX:   3,
		expMaxY:   2,
	}, {
		input:     8,
		expOutput: `M 2 2 h ` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` h -` + deltaStr + ` v ` + deltaStr + ` h -` + deltaStr + ` v -` + deltaStr + ` `,
		expMaxX:   3,
		expMaxY:   2,
	}}

	for _, tc := range testCases {
		actOutput, actMaxX, actMaxY := pb.BuildPath(tc.input)
		assert.Equal(t, tc.expOutput, actOutput, `received unexpected output for "%d"`, tc.input)
		assert.Equal(t, tc.expMaxX, actMaxX, `received unexpected output for "%d"`, tc.input)
		assert.Equal(t, tc.expMaxY, actMaxY, `received unexpected output for "%d"`, tc.input)
	}
}
