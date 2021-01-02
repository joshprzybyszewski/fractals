package drawing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestH(t *testing.T) {
	a := point{
		x: 10,
		y: 20,
	}

	assert.Equal(t, point{
		x: 40,
		y: 20,
	}, a.h(30))

	assert.Equal(t, point{
		x: -20,
		y: 20,
	}, a.h(-30))
}

func TestV(t *testing.T) {
	a := point{
		x: 10,
		y: 20,
	}

	assert.Equal(t, point{
		x: 10,
		y: 50,
	}, a.v(30))

	assert.Equal(t, point{
		x: 10,
		y: -10,
	}, a.v(-30))
}

func TestMaxPoints(t *testing.T) {
	act := maxPoints(
		point{
			x: 10,
			y: 10,
		}, point{
			x: 5,
			y: 20,
		},
	)

	assert.Equal(t, point{
		x: 10,
		y: 20,
	}, act)

	act = maxPoints(
		point{
			x: 10,
			y: 10,
		}, point{
			x: 10,
			y: 10,
		},
	)

	assert.Equal(t, point{
		x: 10,
		y: 10,
	}, act)
}

func TestMinPoints(t *testing.T) {
	act := minPoints(
		point{
			x: 10,
			y: 10,
		}, point{
			x: 5,
			y: 20,
		},
	)

	assert.Equal(t, point{
		x: 5,
		y: 10,
	}, act)

	act = minPoints(
		point{
			x: 10,
			y: 10,
		}, point{
			x: 10,
			y: 10,
		},
	)

	assert.Equal(t, point{
		x: 10,
		y: 10,
	}, act)
}
