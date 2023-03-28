package trendanalysis

import (
	"math"

	"github.com/gonum/stat"
)

func MovingAverage(data []float64, window int) []float64 {
	var smoothed []float64
	for i := 0; i < len(data)-window+1; i++ {
		segment := data[i : i+window]
		average := stat.Mean(segment, nil)
		smoothed = append(smoothed, average)
	}
	for i := len(data) - window + 1; i < len(data); i++ {
		smoothed = append(smoothed, math.NaN())
	}
	return smoothed
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
