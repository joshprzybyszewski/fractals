package drawing

import (
	"github.com/joshprzybyszewski/fractals/generator"
)

const (
	startX = `100`
	startY = `100`
	delta  = `10`
)

func BuildPath(numSteps uint64) string {
	path := `M ` + startX + ` ` + startY + ` `
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
		return `v -` + delta
	case East:
		return `h ` + delta
	case South:
		return `v ` + delta
	case West:
		return `h -` + delta
	}
	return ``
}
