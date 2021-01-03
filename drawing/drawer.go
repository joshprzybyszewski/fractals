package drawing

import (
	"fmt"
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
	// (2 bytes for "v-" + length of the delta string) * how many
	// steps we'll take is a good approximation for the buffer's length
	b.Grow((2 + len(s.deltaStr)) * int(numSteps))

	dir := East

	curPoint := point{}
	minPoint := curPoint
	maxPoint := curPoint

	for step := uint64(1); step <= numSteps; step++ {
		curPoint = s.addToPathBuffer(&b, dir, curPoint)

		checkForLess(&minPoint, curPoint)
		checkForGreater(&maxPoint, curPoint)

		moveDir(&dir, generator.IsLeftTurn(step))
	}

	startX := -minPoint.x + paddingX
	startY := -minPoint.y + paddingY
	initialMove := fmt.Sprintf(
		`M %d %d `,
		startX,
		startY,
	)

	return initialMove + b.String(),
		int64(startX) + maxPoint.x + paddingX,
		int64(startY) + maxPoint.y + paddingY
}

// addToPathBuffer prints the next SVG command for the given Cardinal to the given writer.
// Additionally, the next point on the map is returned based on the given current location.
func (s svgPathBuilder) addToPathBuffer(b *strings.Builder, c Cardinal, curPoint point) point {
	defer b.WriteString(s.deltaStr)

	switch c {
	case North:
		b.WriteString(`v-`)

		return curPoint.v(-s.delta)
	case East:
		b.WriteString(`h`)

		return curPoint.h(s.delta)
	case South:
		b.WriteString(`v`)

		return curPoint.v(s.delta)
	case West:
		b.WriteString(`h-`)

		return curPoint.h(-s.delta)

	default:
		// This is actually an error case
		return point{}
	}
}
