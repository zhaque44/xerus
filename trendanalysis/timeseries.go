package trendanalysis

import (
	"math"

	"github.com/gonum/stat"
	"github.com/pkg/errors"
)

func MovingAverage(data []float64, window int) ([]float64, error) {
	if len(data) < window {
		return nil, errors.New("window size cannot be larger than data length")
	}
	if window <= 0 {
		return nil, errors.New("window size must be greater than zero")
	}
	var smoothed []float64
	for i := 0; i < len(data)-window+1; i++ {
		segment := data[i : i+window]
		average := stat.Mean(segment, nil)
		smoothed = append(smoothed, average)
	}
	for i := len(data) - window + 1; i < len(data); i++ {
		smoothed = append(smoothed, math.NaN())
	}
	return smoothed, nil
}

func MovingAverageMissing(data []float64, window int) []float64 {
	var smoothed []float64
	for i := 0; i < len(data)-window+1; i++ {
		segment := data[i : i+window]
		for j, value := range segment {
			if math.IsNaN(value) {
				// Replace missing value with average of neighboring values
				neighborValues := []float64{}
				if j > 0 {
					neighborValues = append(neighborValues, segment[j-1])
				}
				if j < len(segment)-1 {
					neighborValues = append(neighborValues, segment[j+1])
				}
				if len(neighborValues) > 0 {
					average := stat.Mean(neighborValues, nil)
					segment[j] = average
				}
			}
		}
		average := stat.Mean(segment, nil)
		smoothed = append(smoothed, average)
	}
	for i := len(data) - window + 1; i < len(data); i++ {
		smoothed = append(smoothed, math.NaN())
	}
	return smoothed
}
