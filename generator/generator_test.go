package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsLeftTurn(t *testing.T) {
	testCases := []struct {
		input     uint64
		expOutput bool
	}{{
		input:     0,
		expOutput: true,
	}, {
		input:     1,
		expOutput: true,
	}, {
		input:     2,
		expOutput: true,
	}, {
		input:     3,
		expOutput: false,
	}, {
		input:     4,
		expOutput: true,
	}, {
		input:     5,
		expOutput: true,
	}, {
		input:     6,
		expOutput: false,
	}, {
		input:     7,
		expOutput: false,
	}}

	for _, tc := range testCases {
		actOutput := isLeftTurn(tc.input)
		assert.Equal(t, tc.expOutput, actOutput, `received unexpected output for "%d"`, tc.input)
	}

	power := uint64(2)
	for i := 0; i < 64; i++ {
		power *= 2
		assert.True(t, isLeftTurn(power), `every power of two should be a left turn (%d)`, power)
		assert.False(t, isLeftTurn(power-1), `every number preceding a power of two should be a right turn (%d)`, power-1)
		assert.True(t, isLeftTurn(power+1), `every number after a power of two should be a left turn (%d)`, power+1)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	power := uint64(1)
	assert.True(t, isPowerOfTwo(power))
	power *= 2

	for i := 0; i < 63; i++ {
		power *= 2
		assert.True(t, isPowerOfTwo(power), `should be a power of two because that's what power is (%d)`, power)

		assert.False(t, isPowerOfTwo(power-1), `one fewer than a power is not a power of two (%d)`, power)

		g := power + 1
		if g == 1 {
			break
		}
		assert.False(t, isPowerOfTwo(g), `one greater than a power is not a power of two (%d)`, g)
	}
}

func TestNextPowerOfTwo(t *testing.T) {
	testCases := []struct {
		input     uint64
		expOutput uint64
	}{{
		input:     1,
		expOutput: 1,
	}, {
		input:     2,
		expOutput: 2,
	}, {
		input:     3,
		expOutput: 4,
	}, {
		input:     4,
		expOutput: 4,
	}, {
		input:     5,
		expOutput: 8,
	}, {
		input:     63,
		expOutput: 64,
	}, {
		input:     65,
		expOutput: 128,
	}, {
		input:     111,
		expOutput: 128,
	}, {
		input:     987,
		expOutput: 1024,
	}, {
		input:     32799,
		expOutput: 65536,
	}, {
		input:     7440091,
		expOutput: 8388608,
	}}

	for _, tc := range testCases {
		actOutput := nextPowerOfTwo(tc.input)
		assert.Equal(t, tc.expOutput, actOutput, `unexpected output for input "%d"`, tc.input)
		require.True(t, isPowerOfTwo(actOutput), `output was not a power of two for input "%d"`, tc.input)
	}
}
