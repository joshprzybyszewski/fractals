package drawing

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/joshprzybyszewski/fractals/generator"
)

const (
	// arbitrary paddings chosen to have "some space" around the fractal on screen
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
	// 3 bytes for "v -" + length of the delta string + 1 byte for the following space
	// times how many steps we'll take is a good approximation for the buffer's length
	b.Grow((3 + len(s.deltaStr) + 1) * int(numSteps))

	dir := East

	curPoint := point{}
	minPoint := curPoint
	maxPoint := curPoint

	for step := uint64(1); step <= numSteps; step++ {
		curPoint = s.addToPathBuffer(&b, dir, curPoint)
		fmt.Fprint(&b, ` `)

		minPoint = minPoints(minPoint, curPoint)
		maxPoint = maxPoints(maxPoint, curPoint)

		goLeft := generator.IsLeftTurn(step)
		dir = dir.turn(goLeft)
	}

	initialMove := fmt.Sprintf(`M %d %d`, int(-minPoint.x)+paddingX, int(-minPoint.y)+paddingY)

	return initialMove + b.String(),
		maxPoint.x - minPoint.x + (2 * paddingX),
		maxPoint.y - minPoint.y + (2 * paddingY)
}

// addToPathBuffer prints the next SVG command for the given Cardinal to the given writer.
// Additionally, the next point on the map is returned based on the given current location.
func (s svgPathBuilder) addToPathBuffer(b io.Writer, c Cardinal, curPoint point) point {

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
