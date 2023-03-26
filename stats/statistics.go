package stats

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func CalculateMean(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("Cannot calculate mean of empty data set")
	}
	return stat.Mean(data, nil), nil
}

func CalculateCorrelation(x, y []float64) (float64, error) {
	switch {
	case x == nil || y == nil:
		return 0, fmt.Errorf("Input data sets cannot be nil")
	case len(x) != len(y):
		return 0, fmt.Errorf("Data sets must be of equal length")
	case len(x) < 2 || len(y) < 2:
		return 0, fmt.Errorf("data sets must have at least 2 elements")
	default:
		return stat.Correlation(x, y, nil), nil
	}
}
