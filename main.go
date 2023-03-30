package main

import (
	"fmt"
	"xerus/stats"
)

func main() {

	// MeanPricingModel
	productPrices := []float64{2.49, 3.99, 1.25, 0.99, 2.75}
	competitorPrices := []float64{2.25, 3.75, 1.49, 0.79, 2.99}
	mean, err := stats.MeanPricingModel(productPrices, competitorPrices)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Mean price of product prices with competitor prices: $%.2f\n", mean)

	// 	historicalRevenue := []float64{10000, 12000, 15000, 17000, 20000, 22000, 24000, 26000, 28000, 30000, 32000, 34000}

	// 	// Set the window size for calculating moving average
	// 	windowSize := 3

	// 	// Generate revenue forecast for the next quarter
	// 	forecast, err := trendanalysis.GenerateRevenueForecast(historicalRevenue, windowSize)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate revenue forecast: %v", err)
	// 	}

	// 	// Print the revenue forecast
	// 	fmt.Printf("Revenue forecast for next quarter: $%.2f\n", forecast)
	// }

	// // Test 2: Calculate the mean price of product prices with competitor prices.
	// productPrices = []float64{2.49, 3.99, 1.25, 0.99, 2.75}
	// competitorPrices = []float64{2.25, 3.75, 1.49, 0.79, 2.99}
	// mean, err = stats.CalculateMean(productPrices, competitorPrices)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Mean price of product prices with competitor prices: $%.2f\n", mean)
	// cost := 10000.00
	// leaseTerm := 3
	// interestRate := 0.05
	// residualValue := 0.50
	// maintenanceCost := 500.00
	// repairCost := 20000.00
	// marketDemand := "high"
	// competitiveLandscape := "favorable"
	// taxImplications := "yes"
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
	// output := decisionanalysis.GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	// fmt.Println(output)

	// make Buy scenario
	// input_to_model := decisionanalysis.GetRecommendation(cost, leaseTerm, interestRate, residualValue, &maintenanceCost, &repairCost, &marketDemand, &competitiveLandscape, &taxImplications)
	// fmt.Println(input_to_model)
}
