package pairwise

import (
	"math"

	"github.com/demdxx/ml-gotips/types"
)

// EuclideanDistance (also known as L2 distance).
// D = SQRT( (x1 - ... - z1) ^ 2 + (x2 - ... - z2) ^ 2 )
func EuclideanDistance(v ...types.Vector) (res float64) {
	for i, x := range v[0] {
		for j := 1; j < len(v); j++ {
			if len(v[j]) > i {
				x -= v[j][i]
			}
		}
		res += x * x
	}
	return math.Sqrt(res)
}
