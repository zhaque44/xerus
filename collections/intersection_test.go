package collections

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	assert := assert.New(t)

	// Example usage with slices of different types
	intSlice1 := []int{1, 2, 3, 4, 5}
	intSlice2 := []int{4, 5, 6, 7, 8}
	intersection1, err := intersection(intSlice1, intSlice2)
	assert.NoError(err)
	assert.ElementsMatch(intersection1, []int{4, 5})

	floatSlice1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	floatSlice2 := []float64{4.4, 5.5, 6.6, 7.7, 8.8}
	intersection2, err := intersection(floatSlice1, floatSlice2)
	assert.NoError(err)
	assert.ElementsMatch(intersection2, []float64{4.4, 5.5})

	stringSlice1 := []string{"apple", "banana", "cherry", "date"}
	stringSlice2 := []string{"cherry", "date", "elderberry", "fig"}
	intersection3, err := intersection(stringSlice1, stringSlice2)
	assert.NoError(err)
	assert.ElementsMatch(intersection3, []string{"cherry", "date"})

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
	assert.NoError(err)
	assert.Empty(intersection4, "Training and testing sets should not share data points")
}
