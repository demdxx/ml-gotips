package pairwise

import "math"

// ManhattanDistance function
// D = | x1 - ... - z1 | + | x2 - ... - z2 |
func ManhattanDistance(v ...[]float64) (res float64) {
	for i, x := range v[0] {
		for j := 1; j < len(v); j++ {
			if len(v[j]) > i {
				x -= v[j][i]
			}
		}
		res += math.Abs(x)
	}
	return res
}
