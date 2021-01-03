package drawing

// checkForGreater will update the given max's x and/or y coord if other is greater
func checkForGreater(max *point, other point) {
	if other.x > max.x {
		max.x = other.x
	}
	if other.y > max.y {
		max.y = other.y
	}
}

// checkForLess will update the given min's x and/or y coord if other is less
func checkForLess(min *point, other point) {
	if other.x < min.x {
		min.x = other.x
	}
	if other.y < min.y {
		min.y = other.y
	}
}

// moveDir updates the given Cardinal associated with going left or not.
// That is, if "goLeft" is called on a northernly facing person,
// then c is set to West
func moveDir(c *Cardinal, goLeft bool) {

	if goLeft {
		switch *c {
		case North:
			*c = West
			return
		case East:
			*c = North
			return
		case South:
			*c = East
			return
		case West:
			*c = South
			return
		}
	}

	switch *c {
	case North:
		*c = East
		return
	case East:
		*c = South
		return
	case South:
		*c = West
		return
	case West:
		*c = North
		return
	}
}
