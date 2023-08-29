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

	assert.EqualError(err, "even numbers not allowed")

	// modify the input to include only odd numbers
	input = []int{1, 3, 5}
	expected = []string{"odd", "odd", "odd"}

	result, err = Map(input, transformWithError)
	assert.NoError(err)                    // No error should occur
	assert.NotNil(result)                  // Result should not be nil
	assert.ElementsMatch(result, expected) // Result should match the expected output
}

func TestValues(t *testing.T) {
	assert := assert.New(t)

	// Test case: happy path
	nonEmptyMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	values, err := Values(nonEmptyMap)
	assert.NoError(err)
	assert.ElementsMatch(values, []string{"one", "two", "three"})

	// Test case: Empty map
	emptyMap := make(map[int]string)
	values, err = Values(emptyMap)
	assert.Error(err)
	assert.Nil(values)
	assert.EqualError(err, "map is empty")
}
