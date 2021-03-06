package fractals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoRaised(t *testing.T) {
	assert.Equal(t, uint64(1), twoRaised(0))
	assert.Equal(t, uint64(2), twoRaised(1))
	assert.Equal(t, uint64(4), twoRaised(2))
	assert.Equal(t, uint64(8), twoRaised(3))
	assert.Equal(t, uint64(16), twoRaised(4))
	assert.Equal(t, uint64(32), twoRaised(5))
	assert.Equal(t, uint64(64), twoRaised(6))
	assert.Equal(t, uint64(128), twoRaised(7))
	assert.Equal(t, uint64(256), twoRaised(8))
	assert.Equal(t, uint64(512), twoRaised(9))

	assert.NotEqual(t, uint64(0), twoRaised(63))
	assert.Equal(t, uint64(0), twoRaised(64))
	assert.Equal(t, uint64(0), twoRaised(65))
}
