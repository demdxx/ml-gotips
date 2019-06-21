package pairwise

import (
	"math"

	"github.com/demdxx/ml-gotips/types"
)

// ManhattanDistance function
// D = | x1 - ... - z1 | + | x2 - ... - z2 |
func ManhattanDistance(v ...types.Vector) (res float64) {
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
