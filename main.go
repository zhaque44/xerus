package main

import (
	"fmt"
	"xerus/stats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5}

	stddev, _ := stats.CalculateStdDev(x)

	fmt.Println(stddev)
}
