package pairwise

import (
	"testing"
)

func Test_EuclideanDistance(t *testing.T) {
	var tests = []struct {
		vectors [][]float64
		result  float64
	}{
		{
			vectors: [][]float64{
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
