package collections

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapWithErrors(t *testing.T) {
	assert := assert.New(t)

	// Transformation function that returns an error for even numbers
	transformWithError := func(n int) (string, error) {
		if n%2 == 0 {
			return "", errors.New("even numbers not allowed")
		}
		return "odd", nil
	}

	input := []int{1, 2, 3, 4, 5}
	expected := []string{"odd", "", "odd", "", "odd"}

	result, err := Map(input, transformWithError)
	assert.Error(err)  // Expecting an error due to even numbers
	assert.Nil(result) // Result should be nil due to error

	// Ensure that the error message is as expected
	assert.EqualError(err, "even numbers not allowed")

	// Let's modify the input to include only odd numbers
	input = []int{1, 3, 5}
	expected = []string{"odd", "odd", "odd"}

	result, err = Map(input, transformWithError)
	assert.NoError(err)                    // No error should occur
	assert.NotNil(result)                  // Result should not be nil
	assert.ElementsMatch(result, expected) // Result should match the expected output
}
