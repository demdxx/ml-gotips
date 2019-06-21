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
		if err = v.SetByString(arr[i]); err != nil {
			return nil, err
		}
	}
	return v, nil
}
