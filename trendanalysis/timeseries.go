package trendanalysis

import (
	"math"

	"github.com/gonum/stat"
	"github.com/pkg/errors"
)

func SimpleMovingAverage(data []float64, window int) ([]float64, error) {
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

func SimpleMovingAverageMissing(data []float64, window int) ([]float64, error) {
	if len(data) == 0 {
		return nil, errors.New("input data slice cannot be empty")
	}
	if len(data) < window {
		return nil, errors.New("window size cannot be larger than data length")
	}
	if window <= 0 {
		return nil, errors.New("window size must be greater than zero")
	}
	var hasValidValues bool
	for _, v := range data {
		if !math.IsNaN(v) {
			hasValidValues = true
			break
		}
	}
	if !hasValidValues {
		return nil, errors.New("cannot perform calculation with all NaN values")
	}

	var smoothed []float64
	for i := 0; i < len(data)-window+1; i++ {
		segment := data[i : i+window]
		for j, value := range segment {
			if math.IsNaN(value) {
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
	return smoothed, nil
}

func GenerateRevenueForecast(data []float64, window int) (float64, error) {
	// Check input data and window size
	if len(data) == 0 {
		return 0, errors.New("input data slice cannot be empty")
	}
	if len(data) < window {
		return 0, errors.New("window size cannot be larger than data length")
	}
	if window <= 0 {
		return 0, errors.New("window size must be greater than zero")
	}

	// Preprocess data to replace missing values with neighboring averages
	processedData, err := preprocessData(data)
	if err != nil {
		return 0, err
	}

	// Calculate smoothed moving average
	smoothed, err := calculateSmoothedMovingAverage(processedData, window)
	if err != nil {
		return 0, err
	}

	// Project revenue forecast based on recent trend
	lastValue := smoothed[len(smoothed)-1]
	forecast := lastValue * 1.05

	return forecast, nil
}

// preprocessData replaces missing or NaN values with the average of their neighboring values
func preprocessData(data []float64) ([]float64, error) {
	if len(data) == 0 {
		return nil, errors.New("input data slice cannot be empty")
	}

	processedData := make([]float64, len(data))
	copy(processedData, data)

	for i, value := range processedData {
		if math.IsNaN(value) {
			neighborValues := []float64{}
			if i > 0 {
				neighborValues = append(neighborValues, processedData[i-1])
			}
			if i < len(processedData)-1 {
				neighborValues = append(neighborValues, processedData[i+1])
			}
			if len(neighborValues) > 0 {
				average := stat.Mean(neighborValues, nil)
				processedData[i] = average
			} else {
				return nil, errors.New("cannot preprocess data: too many missing values")
			}
		}
	}

	return processedData, nil
}

func calculateSmoothedMovingAverage(data []float64, window int) ([]float64, error) {
	if len(data) == 0 {
		return nil, errors.New("input data slice cannot be empty")
	}
	if len(data) < window {
		return nil, errors.New("window size cannot be larger than data length")
	}
	if window <= 0 {
		return nil, errors.New("window size must be greater than zero")
	}

	var hasValidValues bool
	for _, v := range data {
		if !math.IsNaN(v) {
			hasValidValues = true
			break
		}
	}
	if !hasValidValues {
		return nil, errors.New("cannot perform calculation with all NaN values")
	}

	smoothed := make([]float64, 0, len(data)-window+1)
	for i := 0; i < len(data)-window+1; i++ {
		windowValues := make([]float64, 0, window)
		for j := i; j < i+window; j++ {
			value := data[j]
			if !math.IsNaN(value) {
				windowValues = append(windowValues, value)
			}
		}
		if len(windowValues) > 0 {
			mean := stat.Mean(windowValues, nil)
			smoothed = append(smoothed, mean)
		} else {
			smoothed = append(smoothed, math.NaN())
		}
	}

	return smoothed, nil
}
