package types

// Vectors list
type Vectors []Vector

// Len return the count of vectors
func (vv Vectors) Len() int {
	return len(vv)
}
