package stats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateStdDev(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}

	stddev, err := CalculateStdDev(x)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.NotNil(t, stddev, "Expected stddev to not be nil")
	assert.NotEqual(t, 0, int(stddev), "Expected stddev not be zero as int")
	assert.NotEqual(t, 0.0, stddev, "Expected stddev not be zero as float64")
}

func TestCalculateMean(t *testing.T) {
	productPrices := []float64{2.49, 3.99, 1.25, 0.99, 2.75}
	competitorPrices := []float64{}
	mean, err := MeanPricingModel(productPrices, competitorPrices)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Mean price of product prices without competitor prices: $%.2f\n", mean)

	// Test 2: Calculate the mean price of product prices with competitor prices.
	productPrices = []float64{2.49, 3.99, 1.25, 0.99, 2.75}
	competitorPrices = []float64{2.25, 3.75, 1.49, 0.79, 2.99}
	mean, err = MeanPricingModel(productPrices, competitorPrices)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Mean price of product prices with competitor prices: $%.2f\n", mean)
}

func TestCalculateVariance(t *testing.T) {
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	variance, err := CalculateVariance(values)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedVariance := 2.5
	if variance != expectedVariance {
		t.Errorf("Expected variance: %f, but got: %f", expectedVariance, variance)
	}
}
