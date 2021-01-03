package fractals

import (
	"fmt"

	"github.com/joshprzybyszewski/fractals/drawing"
)

func twoRaised(n uint64) uint64 {
	return uint64(1) << n
}

func getPathAndViewBoxForDragon(n uint64) (string, string) {
	// get the SVG path
	path, maxX, maxY := drawing.New(2).BuildPath(twoRaised(n))

	// get the viewbox string set
	vb := fmt.Sprintf("0 0 %d %d", maxX, maxY)

	return path, vb
}
