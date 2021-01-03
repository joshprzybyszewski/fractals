package generator

var (
	useCache      = false
	leftTurnCache = make(map[uint64]bool)
)

// IsLeftTurn returns true if the next turn for the path is left-handed
// n is the number of path segments we've currently walked
// The first turn is left handed. So is the second.
// And every time we have walked 2^x segments, we know we will make
// a left-handed turn because that is how the fractal repeats itself.
// Otherwise, we think of building the "next fold" as walking the previous fold
// backwards. This means that we think of the value for `seg(n)` as the inverse
// value of `seg(l - (n-l))` where `l` is the power of two < n. In other words,
// `seg(n)` = `seg(g -n)` where `g` = 2^x > n > 2^(x-1) = `l`.
//
// This means that each call to `IsLeftTurn` won't have more than log2(n)
// function calls on the stack. I chose this implementation because I wanted
// to reduce the size of memory requiring to build a path. This means that
// path building is a O(nlogn) algorithm. We could reduce the run-time complexity
// of path building to O(n) by keeping a map of calculated values, but that
// would require O(n) memory (instead of the currently required O(logn) ).
func IsLeftTurn(n uint64) (val bool) {

	// As it turns out, benchmarks show that using a cache makes the func much slower
	// at higher inputs of `n`. I suspect this is because the memory required
	// to keep the cache isn't efficient to look-up from disk as compared to just
	// recursing down a stack
	if useCache {
		if cachedVal, ok := leftTurnCache[n]; ok {
			return cachedVal
		}
		defer func() {
			leftTurnCache[n] = val
		}()
	}

	if isPowerOfTwo(n) {
		return true
	}

	// g is the next power of two greater than n
	// such that g > n && g/2 < n
	g := nextPowerOfTwo(n)

	return !IsLeftTurn(g - n)
}

func isPowerOfTwo(n uint64) bool {
	// This solution was found at:
	// https://stackoverflow.com/questions/600293/how-to-check-if-a-number-is-a-power-of-2
	return (n & (n - 1)) == 0
}

// nextPowerOfTwo returns the next power of two
// That is, this returns g = 2^x > n > 2^(x-1).
// Since there will only be 64 powers of two max, we could potentially
// speed this up by keeping the calculated slice of powers of two in memory
// and just iterating that slice. However, that takes O(logn) time, and this takes
// constant time O(1).
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
