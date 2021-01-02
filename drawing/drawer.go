package drawing

import (
	"github.com/joshprzybyszewski/fractals/generator"
)

func buildPath(numSteps uint64) string {
	path := `M 0 0 `
	dir := East

	for step := uint64(1); step <= numSteps; step++ {
		path += toPath(dir) + ` `

		goLeft := generator.IsLeftTurn(step)
		dir = dir.transform(goLeft)
	}

	return path
}

func toPath(c Cardinal) string {
	switch c {
	case North:
		return `l 0 -1`
	case East:
		return `l 1 0`
	case South:
		return `l 0 1`
	case West:
		return `l -1 0`
	}
	return ``
}
