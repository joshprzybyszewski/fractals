package drawing

import (
	"strconv"

	"github.com/joshprzybyszewski/fractals/generator"
)

const (
	paddingX = 5
	paddingY = 5
)

type svgPathBuilder struct {
	delta    int
	deltaStr string
}

func New(delta int) PathBuilder {
	return svgPathBuilder{
		delta:    delta,
		deltaStr: strconv.Itoa(delta),
	}
}

func (s svgPathBuilder) BuildPath(numSteps uint64) (string, int64, int64) {
	// TODO come up with a better way to build a string
	path := ``
	dir := East

	curPoint := point{}
	minPoint := curPoint
	maxPoint := curPoint

	for step := uint64(1); step <= numSteps; step++ {
		var pDelta string
		pDelta, curPoint = s.nextPath(dir, curPoint)
		path += pDelta + ` `

		minPoint = minPoints(minPoint, curPoint)
		maxPoint = maxPoints(maxPoint, curPoint)

		goLeft := generator.IsLeftTurn(step)
		dir = dir.transform(goLeft)
	}

	initialMove := `M ` + strconv.Itoa(int(-minPoint.x)+paddingX) + ` ` + strconv.Itoa(int(-minPoint.y)+paddingY) + ` `

	return initialMove + path, maxPoint.x - minPoint.x + (2 * paddingX), maxPoint.y - minPoint.y + (2 * paddingY)
}

func (s svgPathBuilder) nextPath(c Cardinal, curPoint point) (string, point) {

	switch c {
	case North:
		return `v -` + s.deltaStr, curPoint.v(-s.delta)
	case East:
		return `h ` + s.deltaStr, curPoint.h(s.delta)
	case South:
		return `v ` + s.deltaStr, curPoint.v(s.delta)
	case West:
		return `h -` + s.deltaStr, curPoint.h(-s.delta)
	}

	return ``, curPoint
}
