package main

// Import my stat-analysis file

import (
	"fmt"
	"xerus/stats"
)

func main() {
	x := []float64{1}
	y := []float64{1}

	// call the function, stat-analysis.go called CalculateCorrelation
	corr, err := stats.CalculateCorrelation(x, y)
	// print the result
	fmt.Println(corr, err)
}
