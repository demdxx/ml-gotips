package knn

import (
	"testing"
)

func Test_FitAndPredict(t *testing.T) {
	knn := NewClassifier()
	knn.Fit()
}
