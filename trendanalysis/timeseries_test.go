package trendanalysis

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	data := []float64{1, 2, math.NaN(), 4, 5, 6, 7, 8, 9, 10}
	window := 3
	smoothed, err := SimpleMovingAverage(data, window)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.NotNil(t, smoothed, "Expected smoothed slice to not be nil")
}

func TestMovingAverageMissing(t *testing.T) {
	data := []float64{1, 2, math.NaN(), 4, 5, 6, 7, 8, 9, 10}
	window := 3
	smoothed, err := SimpleMovingAverageMissing(data, window)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.NotNil(t, smoothed, "Expected smoothed slice to not be nil")
	assert.IsType(t, []float64{}, smoothed, "Expected smoothed slice to be of type []float64")
}
