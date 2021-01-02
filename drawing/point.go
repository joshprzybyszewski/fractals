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

// maxPoints returns a new point that represents the max (x,y) value between the input points
func maxPoints(max, other point) point {
	if other.x > max.x {
		max.x = other.x
	}
	if other.y > max.y {
		max.y = other.y
	}
	return max
}

// minPoints returns a new point that represents the minimum (x,y) value between the input points
func minPoints(min, other point) point {
	if other.x < min.x {
		min.x = other.x
	}
	if other.y < min.y {
		min.y = other.y
	}
	return min
}
