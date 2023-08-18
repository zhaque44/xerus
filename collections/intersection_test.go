package collections

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIntersection(t *testing.T) {
	// Example usage with slices of different types
	intSlice1 := []int{1, 2, 3, 4, 5}
	intSlice2 := []int{4, 5, 6, 7, 8}
	intersection1, _ := intersection(intSlice1, intSlice2)
	fmt.Println("Intersection of int slices:", intersection1)

	floatSlice1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	floatSlice2 := []float64{4.4, 5.5, 6.6, 7.7, 8.8}
	intersection2, _ := intersection(floatSlice1, floatSlice2)
	fmt.Println("Intersection of float slices:", intersection2)

	stringSlice1 := []string{"apple", "banana", "cherry", "date"}
	stringSlice2 := []string{"cherry", "date", "elderberry", "fig"}
	intersection3, _ := intersection(stringSlice1, stringSlice2)
	fmt.Println("Intersection of string slices:", intersection3)

	// Generate a sample dataset of custom type
	type DataPoint struct {
		X float64
		Y float64
	}
	data := make([]DataPoint, 100)
	for i := range data {
		data[i] = DataPoint{rand.Float64(), rand.Float64()}
	}

	// Split the dataset into training and testing sets
	splitIndex := len(data) * 70 / 100 // 70% training, 30% testing
	training := data[:splitIndex]
	testing := data[splitIndex:]

	// Ensure there are no common data points between the two sets
	intersection4, err := intersection(training, testing)
	if err != nil {
		fmt.Println("Error:", err)
	} else if len(intersection4) > 0 {
		fmt.Println("Error: training and testing sets share data points")
	} else {
		fmt.Println("Training and testing sets do not share data points")
	}
}
