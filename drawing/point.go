package drawing

type point struct {
	x, y int64
}

// h returns a point that is horizontal to the current point by the given delta
func (p point) h(delta int) point {
	return point{
		x: p.x + int64(delta),
		y: p.y,
	}
}

// h returns a point that is vertical to the current point by the given delta
func (p point) v(delta int) point {
	return point{
		x: p.x,
		y: p.y + int64(delta),
	}
}
