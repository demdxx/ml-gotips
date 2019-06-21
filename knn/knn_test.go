package knn

import (
	"fmt"
	"testing"

	"github.com/demdxx/ml-gotips/metrics/pairwise"
	"github.com/demdxx/ml-gotips/utils/converters"
	"github.com/demdxx/ml-gotips/utils/csv"
	// "github.com/demdxx/ml-gotips/types"
)

func Test_FitAndPredict(t *testing.T) {
	knn := NewClassifier(pairwise.EuclideanDistanceType)
	grid, err := csv.ReadAll("../datasets/iris.csv")
	if err != nil {
		t.Errorf("Invalid loading dataset: %s", err)
		return
	}

	if len(grid) < 1 {
		t.Errorf("Datatset is empty")
		return
	}

	headers := grid[0]
	params, _ := converters.StringGridToVectors(grid[1:], 0, 4)

	// knn.Fit()
	fmt.Println(">>> grid", headers, params, knn)
}
