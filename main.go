package main

import (
	"fmt"

	"xerus/decisionanalysis"
)

func main() {
	// Test 1: Basic usage
	// Inputs: 40000, 36, 0.03, 15000, nil, nil, nil, nil, nil
	// Expected output: "Lease"
	output := decisionanalysis.GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	fmt.Println(output)

	// Test 2: With maintenance cost and repair cost
	// Inputs: 35000, 48, 0.04, 10000, 500, 800, nil, nil, nil
	// Expected output: "Buy"
	// maintenanceCost := 500.0
	// repairCost := 800.0
	// output = decisionanalysis.GetRecommendation(35000, 48, 0.04, 10000, &maintenanceCost, &repairCost, nil, nil, nil)
	// fmt.Printf("Test 2: %v\n", output == "Buy")

	// Test 3: With market demand and competitive landscape
	// Inputs: 30000, 24, 0.05, 8000, nil, nil, "High", "Very Competitive", nil
	// Expected output: "Lease"
	// marketDemand := "High"
	// competitiveLandscape := "Very Competitive"
	// output = decisionanalysis.GetRecommendation(30000, 24, 0.05, 8000, nil, nil, &marketDemand, &competitiveLandscape, nil)
	// fmt.Printf("Test 3: %v\n", output == "Lease")

	// Test 4: With tax implications
	// Inputs: 45000, 36, 0.03, 20000, nil, nil, nil, nil, "Positive"
	// Expected output: "Lease"
	// taxImplications := "Positive"
	// output = decisionanalysis.GetRecommendation(45000, 36, 0.03, 20000, nil, nil, nil, nil, &taxImplications)
	// fmt.Printf("Test 4: %v\n", output == "Lease")

	// Test 5: With all inputs
	// Inputs: 40000, 48, 0.06, 12000, 1000, 1500, "Low", "Moderately Competitive", "Negative"
	// Expected output: "Buy"
	// maintenanceCost = 1000.0
	// repairCost = 1500.0
	// marketDemand = "Low"
	// competitiveLandscape = "Moderately Competitive"
	// taxImplications = "Negative"
	// output = decisionanalysis.GetRecommendation(40000, 48, 0.06, 12000, &maintenanceCost, &repairCost, &marketDemand, &competitiveLandscape, &taxImplications)
	// fmt.Printf("Test 5: %v\n", output == "Buy")
}
