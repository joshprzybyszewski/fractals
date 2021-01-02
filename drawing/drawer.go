package drawing

import (
	"fmt"
	"io"
	"strconv"
	"strings"

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
	var b strings.Builder
	b.Grow((3 + len(s.deltaStr) + 1) * int(numSteps))

	dir := East

	curPoint := point{}
	minPoint := curPoint
	maxPoint := curPoint

	for step := uint64(1); step <= numSteps; step++ {
		curPoint = s.nextPath(&b, dir, curPoint)
		fmt.Fprint(&b, ` `)

		minPoint = minPoints(minPoint, curPoint)
		maxPoint = maxPoints(maxPoint, curPoint)

		goLeft := generator.IsLeftTurn(step)
		dir = dir.transform(goLeft)
	}

	initialMove := `M ` + strconv.Itoa(int(-minPoint.x)+paddingX) + ` ` + strconv.Itoa(int(-minPoint.y)+paddingY) + ` `

	return initialMove + b.String(), maxPoint.x - minPoint.x + (2 * paddingX), maxPoint.y - minPoint.y + (2 * paddingY)
}

func (s svgPathBuilder) nextPath(b io.Writer, c Cardinal, curPoint point) point {

	switch c {
	case North:
		fmt.Fprint(b, `v -`)
		fmt.Fprint(b, s.deltaStr)
		return curPoint.v(-s.delta)
	case East:
		fmt.Fprint(b, `h `)
		fmt.Fprint(b, s.deltaStr)
		return curPoint.h(s.delta)
	case South:
		fmt.Fprint(b, `v `)
		fmt.Fprint(b, s.deltaStr)
		return curPoint.v(s.delta)
	case West:
		fmt.Fprint(b, `h -`)
		fmt.Fprint(b, s.deltaStr)
		return curPoint.h(-s.delta)
	}

	return curPoint
}
