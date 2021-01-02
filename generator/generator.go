package generator

func isLeftTurn(n uint64) bool {
	if isPowerOfTwo(n) {
		return true
	}

	// g is the next power of two greater than n
	// such that g > n && g/2 < n
	g := nextPowerOfTwo(n)

	return !isLeftTurn(g - n)
}

func isPowerOfTwo(n uint64) bool {
	// This solution was found at:
	// https://stackoverflow.com/questions/600293/how-to-check-if-a-number-is-a-power-of-2
	return (n & (n - 1)) == 0
}

func nextPowerOfTwo(n uint64) uint64 {
	/*
		This C answer is from http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
		unsigned int v; // compute the next highest power of 2 of 32-bit v

		v--;
		v |= v >> 1;
		v |= v >> 2;
		v |= v >> 4;
		v |= v >> 8;
		v |= v >> 16;
		v++;
	*/

	n -= 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n += 1

	return n
}
