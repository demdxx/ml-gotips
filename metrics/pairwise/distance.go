package pairwise

import (
	"github.com/demdxx/ml-gotips/types"
)

type (
	DistanceType int
	DistanceFnk  func(v ...types.Vector) float64
)

const (
	UndefinedDistanceType DistanceType = 0
	EuclideanDistanceType DistanceType = 1
	ManhattanDistanceType DistanceType = 2
)

// Fnk returns current distance function
// by default it's Euclidean
func (d DistanceType) Fnk() DistanceFnk {
	switch d {
	case ManhattanDistanceType:
		return ManhattanDistance
	}
	return EuclideanDistance
}
