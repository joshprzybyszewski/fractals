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
