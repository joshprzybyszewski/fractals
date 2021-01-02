package drawing

type point struct {
	x, y int64
}

func (p point) h(delta int) point {
	return point{
		x: p.x + int64(delta),
		y: p.y,
	}
}

func (p point) v(delta int) point {
	return point{
		x: p.x,
		y: p.y + int64(delta),
	}
}

func maxPoints(max, other point) point {
	if other.x > max.x {
		max.x = other.x
	}
	if other.y > max.y {
		max.y = other.y
	}
	return max
}

func minPoints(min, other point) point {
	if other.x < min.x {
		min.x = other.x
	}
	if other.y < min.y {
		min.y = other.y
	}
	return min
}
