package drawing

type Cardinal int

const (
	Unknown Cardinal = 0
	North   Cardinal = 1
	East    Cardinal = 2
	South   Cardinal = 3
	West    Cardinal = 4
)

// turn returns the Cardinal associated with going left or not.
// That is, if "goLeft" is called on a northernly facing person,
// then west is returned
func (c Cardinal) turn(goLeft bool) Cardinal {
	if goLeft {
		switch c {
		case North:
			return West
		case East:
			return North
		case South:
			return East
		case West:
			return South
		}
	}

	switch c {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}

	// Realistically, this will never be hit
	return Unknown
}
