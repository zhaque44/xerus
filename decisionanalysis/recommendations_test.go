package decisionanalysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyingRecommendation(t *testing.T) {
	output := GetRecommendation(40000, 36, 0.03, 15000, nil, nil, nil, nil, nil)
	assert.Equal(t, "Buying is recommended.", output, "Buying is recommended.")
}

// func TestGetRecommendation(t *testing.T) {
// 	// Set up test inputs
// 	cost := 10000.00
// 	leaseTerm := 3
// 	interestRate := 0.05
// 	residualValue := 0.50
// 	maintenanceCost := 500.00
// 	repairCost := 0.00
// 	marketDemand := "high"
// 	competitiveLandscape := "favorable"
// 	taxImplications := "yes"

// 	// Create temporary file for input
// 	tmpfile, err := ioutil.TempFile("", "example")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.Remove(tmpfile.Name()) // clean up

// 	// Ask user for input
// 	fmt.Fprintf(tmpfile, "y\ny\ny\ny\ny\n")
// 	if _, err := tmpfile.Seek(0, 0); err != nil {
// 		t.Fatal(err)
// 	}
// 	os.Stdin = tmpfile

// 	// Call the function and check the result
// 	result := GetRecommendation(cost, leaseTerm, interestRate, residualValue, &maintenanceCost, &repairCost, &marketDemand, &competitiveLandscape, &taxImplications)
// 	expected := "Leasing is recommended."
// 	if result != expected {
// 		t.Errorf("GetRecommendation() returned %q, expected %q", result, expected)
// 	}

// 	// Test with market demand set to nil
// 	var marketDemandNil *string = nil
// 	fmt.Fprintf(tmpfile, "n\n")
// 	if _, err := tmpfile.Seek(0, 0); err != nil {
// 		t.Fatal(err)
// 	}
// 	os.Stdin = tmpfile

// 	result = GetRecommendation(cost, leaseTerm, interestRate, residualValue, &maintenanceCost, &repairCost, marketDemandNil, &competitiveLandscape, &taxImplications)
// 	expected = "Buying is recommended."
// 	if result != expected {
// 		t.Errorf("GetRecommendation() returned %q, expected %q", result, expected)
// 	}
// }
