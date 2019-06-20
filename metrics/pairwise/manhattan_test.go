package pairwise

import (
	"testing"
)

func Test_ManhattanDistance(t *testing.T) {
	var tests = []struct {
		vectors [][]float64
		result  float64
	}{
		{
			vectors: [][]float64{
				{1, 2, 3},
				{2, 4, 5},
			},
			result: 5,
		},
	}

	for _, test := range tests {
		if distance := ManhattanDistance(test.vectors...); distance != test.result {
			t.Errorf("Invalid distance %f, must be %f", distance, test.result)
		}
	}
}
