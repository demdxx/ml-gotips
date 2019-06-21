package pairwise

import (
	"testing"

	"github.com/demdxx/ml-gotips/types"
)

func Test_EuclideanDistance(t *testing.T) {
	var tests = []struct {
		vectors types.Vectors
		result  float64
	}{
		{
			vectors: types.Vectors{
				{1, 2, 3},
				{2, 4, 5},
			},
			result: 3,
		},
	}

	for _, test := range tests {
		if distance := EuclideanDistance(test.vectors...); distance != test.result {
			t.Errorf("Invalid distance %f, must be %f", distance, test.result)
		}
	}
}
