package converters

import (
	"github.com/demdxx/ml-gotips/types"
)

// StringArrayToVector converter
func StringArrayToVector(arr []string, from, to int) (_ types.Vector, err error) {
	v := types.NewVector(to - from)
	for i := from; i < len(arr); i++ {
		if i >= to {
			break
		}
		if err = v.SetByString(i-from, arr[i]); err != nil {
			return nil, err
		}
	}
	return v, nil
}

// StringGridToVectors converts two dimentional string array to the vector table
func StringGridToVectors(grid [][]string, from, to int) (_ types.Vectors, err error) {
	vectors := make(types.Vectors, len(grid))
	for i, row := range grid {
		vectors[i], err = StringArrayToVector(row, from, to)
		if err != nil {
			return nil, err
		}
	}
	return vectors, nil
}
