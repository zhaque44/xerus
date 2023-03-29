package main

import (
	"fmt"

	"xerus/decisionanalysis"
)

func main() {
	// Test 1: Basic usage
	// Inputs: 40000, 36, 0.03, 15000, nil, nil, nil, nil, nil
	// Expected output: "Lease"
	// output := decisionanalysis.GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	// fmt.Println(output)

	// // Test 2: With maintenance cost
	// // Inputs: 35000, 48, 0.04, 10000, 500, nil, nil, nil, nil
	// // Expected output: "Buy"
	// maintenanceCost := 500.0
	// o := decisionanalysis.GetRecommendation(35000, 48, 0.04, 10000, &maintenanceCost, nil, nil, nil, nil)
	// fmt.Println(o)

	// make Lease scenario
	// Inputs: 40000, 36, 0.03, 15000, nil, nil, nil, nil, nil
	// Expected output: "Lease"
	// mnt := 1500.0
	// repair := 500.0
	output := decisionanalysis.GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	fmt.Println(output)
}
