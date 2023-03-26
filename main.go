package main

import (
	"fmt"
	"xerus/stats"
)

func main() {
	// run the covariance function

	// create two slices of float64
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 2, 3, 4, 5}

	// call the function
	cov, _ := stats.CalculateCovariance(x, y)

	// print the result
	fmt.Println(cov, "is the covariance of x and y")
}
