package collections

import (
	"errors"
	"fmt"
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

func TestContainsKey(t *testing.T) {
	assert := assert.New(t)

	// Test case: Key exists in the map
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	key := 2
	exists := ContainsKey(m, key)
	assert.True(exists)

	// Test case: Key doesn't exist in the map
	key = 4
	exists = ContainsKey(m, key)
	assert.False(exists)

	// Additional test case: Key exists in the map with different type
	mString := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	keyString := "two"
	existsString := ContainsKey(mString, keyString)
	assert.True(existsString)

	// Additional test case: Key doesn't exist in the map with different type
	keyString = "four"
	existsString = ContainsKey(mString, keyString)
	assert.False(existsString)
}

func TestFoldLeft(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	sum := foldLeft(slice1, slice2, 0, func(result, elem1, elem2 int) int {
		return result + elem1 + elem2
	})
	fmt.Println(sum) // 21
}

func TestFoldRight(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	sum := foldRight(slice1, slice2, 0, func(result, elem1, elem2 int) int {
		return result + elem1 + elem2
	})
	fmt.Println(sum) // 21
}

func TestEmptyMap(t *testing.T) {
	assert := assert.New(t)

	// Test case: Empty map
	m := make(map[string]string)
	empty := emptyMap(m)
	assert.True(empty)

	// Test case: Non-empty map
	m["foo"] = "bar"
	empty = emptyMap(m)
	assert.False(empty)
}
