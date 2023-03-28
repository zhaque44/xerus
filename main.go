package main

import (
	"fmt"
	"math"

	"xerus/trendanalysis"
)

func main() {
	// Example usage of the MovingAverage and MovingAverageMissing functions
	data := []float64{1, 2, math.NaN(), 4, 5, 6, 7, 8, 9, 10}
	window := 3
	smoothed := trendanalysis.MovingAverage(data, window)
	fmt.Println(smoothed)
	smoothedMissing := trendanalysis.MovingAverageMissing(data, window)
	fmt.Println(smoothedMissing)

}
