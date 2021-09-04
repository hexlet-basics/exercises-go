package solution

import "math"

// BEGIN

// MinInt returns min int from x and y
func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

// END
