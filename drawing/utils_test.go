package drawing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForGreater(t *testing.T) {
	max := point{
		x: 10,
		y: 10,
	}
	checkForGreater(&max, point{
		x: 5,
		y: 20,
	})

	assert.Equal(t, point{
		x: 10,
		y: 20,
	}, max)

	max = point{
		x: 10,
		y: 10,
	}
	checkForGreater(&max, point{
		x: 10,
		y: 10,
	})
	assert.Equal(t, point{
		x: 10,
		y: 10,
	}, max)
}

func TestCheckForLess(t *testing.T) {
	min := point{
		x: 10,
		y: 10,
	}
	checkForLess(&min, point{
		x: 5,
		y: 20,
	})

	assert.Equal(t, point{
		x: 5,
		y: 10,
	}, min)

	min = point{
		x: 10,
		y: 10,
	}
	checkForLess(&min, point{
		x: 10,
		y: 10,
	})
	assert.Equal(t, point{
		x: 10,
		y: 10,
	}, min)
}

func TestMoveDir(t *testing.T) {
	assert.Panics(t, func() {
		moveDir(nil, true)
	})
	assert.Panics(t, func() {
		moveDir(nil, false)
	})

	testCases := []struct {
		input  Cardinal
		goLeft bool
		exp    Cardinal
	}{{
		input:  North,
		goLeft: true,
		exp:    West,
	}, {
		input:  South,
		goLeft: true,
		exp:    East,
	}, {
		input:  East,
		goLeft: true,
		exp:    North,
	}, {
		input:  West,
		goLeft: true,
		exp:    South,
	}, {
		input:  North,
		goLeft: false,
		exp:    East,
	}, {
		input:  South,
		goLeft: false,
		exp:    West,
	}, {
		input:  East,
		goLeft: false,
		exp:    South,
	}, {
		input:  West,
		goLeft: false,
		exp:    North,
	}}

	for _, tc := range testCases {
		c := tc.input
		moveDir(&c, tc.goLeft)
		assert.Equal(t, tc.exp, c)
	}
}
