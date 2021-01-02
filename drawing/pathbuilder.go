package drawing

type PathBuilder interface {
	BuildPath(numSteps uint64) (path string, maxX, maxY int64)
}
