package knn

import (
	"fmt"
	"sort"

	"github.com/demdxx/ml-gotips/metrics/pairwise"
	"github.com/demdxx/ml-gotips/types"
)

// Classifier KNN object
type Classifier struct {
	distance pairwise.DistanceType
}

func NewClassifier(distance pairwise.DistanceType) *Classifier {
	return &Classifier{distance: distance}
}

// Fit prepare and train dataset
func (knn *Classifier) Fit(trainingSet, testingSet types.Vectors) error {
	var (
		distances   = make([]float64, trainingSet.Len())
		distanceFnk = knn.distance.Fnk()
	)

	for i, v := range trainingSet {
		distances[i] = distanceFnk(v, testingSet[i])
	}

	sort.Float64s(distances)
	fmt.Println("distances", distances)

	return nil
}
