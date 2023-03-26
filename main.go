package main

import (
	"fmt"
	"xerus/stats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5}

	percentile, _ := stats.CalculatePercentile(x, 0.5)

	fmt.Println(percentile)
}
