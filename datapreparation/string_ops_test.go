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
