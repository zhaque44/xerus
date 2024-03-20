package datapreparation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsPrefix(t *testing.T) {
	// Test case 1: Input string contains one of the prefixes
	input := "Hello"
	prefixes := []string{"hello", "foo", "bar"}
	expected := true

	result := ContainsPrefix(input, prefixes...)
	assert.Equal(t, expected, result)

	// Test case 2: Input string does not contain any of the prefixes
	input = "Goodbye"
	prefixes = []string{"hello", "foo", "bar"}
	expected = false

	result = ContainsPrefix(input, prefixes...)
	assert.Equal(t, expected, result)

	// Test case 3: Empty input string should always return false
	input = ""
	prefixes = []string{"hello", "foo", "bar"}
	expected = false

	result = ContainsPrefix(input, prefixes...)
	assert.Equal(t, expected, result)
}

func TestSplitAndTrimWhitespace(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		inputString := "apple,   banana, cherry,   date"
		splitchar := ","
		expected := []string{"apple", "banana", "cherry", "date"}

		result, err := SplitAndTrimWhitespace(inputString, splitchar)

		// Assert that there are no errors
		assert.Nil(t, err)

		// Assert the expected output
		assert.Equal(t, expected, result)
	})

	t.Run("Empty splitchar", func(t *testing.T) {
		inputString := "apple, banana, cherry, date"
		splitchar := ""
		_, err := SplitAndTrimWhitespace(inputString, splitchar)

		// Assert that an error is returned
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "splitchar cannot be empty")
	})
}

func TestConverter(t *testing.T) {
	// Test input data
	jsonData := `{"key": "value"}`
	urlStr := "http://example.com/path"

	// Expected output for JSON data
	expectedJSONData := map[string]interface{}{"key": "value"}

	// Expected output for URL parts
	expectedURLParts := []string{"http://example.com/path", "path"}

	// Test JSON data conversion
	jsonResult, err := converter(jsonData, "")
	assert.NoError(t, err, "Unexpected error during JSON data conversion")
	assert.Equal(t, expectedJSONData, jsonResult, "JSON data conversion failed")

	// Test URL conversion
	urlResult, err := converter("", urlStr)
	assert.NoError(t, err, "Unexpected error during URL conversion")
	assert.Equal(t, expectedURLParts, urlResult, "URL conversion failed")
}
