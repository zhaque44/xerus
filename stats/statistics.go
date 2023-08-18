package stats

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

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

func CalculateCovariance(x, y []float64) (float64, error) {
	switch {
	case x == nil || y == nil:
		return 0, fmt.Errorf("input data sets cannot be nil")
	case len(x) != len(y):
		return 0, fmt.Errorf("data sets must be of equal length")
	case len(x) < 2 || len(y) < 2:
		return 0, fmt.Errorf("data sets must have at least 2 elements")
	default:
		return stat.Covariance(x, y, nil), nil
	}
}

func MeanPricingModel(productPrices []float64, competitorPrices []float64) (float64, error) {
	if len(productPrices) == 0 {
		return 0, fmt.Errorf("Cannot calculate mean of empty data set")
	}

	if len(competitorPrices) > 0 {
		if len(productPrices) != len(competitorPrices) {
			return 0, fmt.Errorf("Product prices and competitor prices must have the same length")
		}

		var weightedPrices []float64
		for i, price := range productPrices {
			competitorPrice := competitorPrices[i]
			weightedPrices = append(weightedPrices, 0.5*price+0.5*competitorPrice)
		}
		return stat.Mean(weightedPrices, nil), nil
	} else {
		return stat.Mean(productPrices, nil), nil
	}
}

func CalculatePercentile(values []float64, p float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("input values cannot be empty")
	}

	isInt := true
	for _, v := range values {
		if v != math.Floor(v) {
			isInt = false
			break
		}
	}

	if isInt {
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	} else {
		sort.Float64s(values)
	}

	n := float64(len(values))
	index := (n-1)*p/100.0 + 1

	if index == 1 {
		return values[0], nil
	} else if index == n {
		return values[len(values)-1], nil
	} else {
		k := int(index)
		d := index - float64(k)
		return (1-d)*values[k-1] + d*values[k], nil
	}
}

func CalculateStdDev(values []float64) (float64, error) {
	if values == nil || len(values) == 0 {
		return 0, fmt.Errorf("input values cannot be empty")
	}

	mean := floats.Sum(values) / float64(len(values))

	sumSqDiff := 0.0
	for _, x := range values {
		diff := x - mean
		sumSqDiff += diff * diff
	}

	variance := sumSqDiff / float64(len(values)-1)

	return math.Sqrt(variance), nil
}

func CalculateVariance(values []float64) (float64, error) {
	if values == nil || len(values) == 0 {
		return 0, errors.New("input values cannot be empty")
	}

	mean := stat.Mean(values, nil)

	sumSqDiff := 0.0
	for _, x := range values {
		diff := x - mean
		sumSqDiff += diff * diff
	}

	return sumSqDiff / float64(len(values)-1), nil
}

// func CalculateZScore(value float64, values []float64) (float64, error) {
// 	if values == nil || len(values) == 0 {
// 		return 0, fmt.Errorf("input values cannot be empty")
// 	}

// 	mean := floats.Sum(values) / float64(len(values))
// 	stdDev, err := CalculateStdDev(values)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return (value - mean) / stdDev, nil
// }

// func CalculateZScoreForAll(values []float64) ([]float64, error) {
// 	if values == nil || len(values) == 0 {
// 		return nil, fmt.Errorf("input values cannot be empty")
// 	}

// 	mean := floats.Sum(values) / float64(len(values))
// 	stdDev, err := CalculateStdDev(values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var zScores []float64
// 	for _, value := range values {
// 		zScores = append(zScores, (value-mean)/stdDev)
// 	}

// 	return zScores, nil
// }

// func CalculateAltmanZScore(x1, x2, x3, x4, x5 float64) (float64, error) {
// 	if x1 == 0 || x2 == 0 || x3 == 0 || x4 == 0 || x5 == 0 {
// 		return 0, fmt.Errorf("input values cannot be 0")
// 	}

// 	return 1.2*x1 + 1.4*x2 + 3.3*x3 + 0.6*x4 + 1.0*x5, nil
// }
