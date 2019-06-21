package types

import (
	"fmt"
	"strconv"
)

// Vector of float values
type Vector []float64

// NewVector with specific size
func NewVector(size int) Vector {
	return make(Vector, size)
}

// Get value from the vector or return 0
func (v Vector) Get(i int) float64 {
	if v.Len() <= i {
		return v[i]
	}
	return 0
}

// Set value of the vector
func (v Vector) Set(i int, vl float64) error {
	if len(v) <= i {
		return fmt.Errorf("[vector] invalid vector index %d", i)
	}
	v[i] = vl
	return nil
}

// SetByString value parsing
func (v Vector) SetByString(i int, vl string) error {
	flVl, err := strconv.ParseFloat(vl, 64)
	if err != nil {
		return fmt.Errorf("[vector] parse float error: %s", err)
	}
	return v.Set(i, flVl)
}

// Len of the vector
func (v Vector) Len() int {
	return len(v)
}

// Sub vector and return new vector value
func (v Vector) Sub(v2 Vector) Vector {
	res := NewVector(v.Len())
	return res.SubApply(v2)
}

// SubApply of the current vector
func (v Vector) SubApply(v2 Vector) Vector {
	for i := 0; i < v.Len(); i++ {
		v[i] -= v2.Get(i)
	}
	return v
}
