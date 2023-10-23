package datapreparation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeURL(t *testing.T) {
	inputURL := "http://example.com"
	defaultPort := true

	expectedURL := "http://example.com:80"

	result, err := SanitizeURL(inputURL, defaultPort)

	assert.Equal(t, expectedURL, result)
	assert.Nil(t, err)
}

func TestSanitizeURLWithoutDefaultPort(t *testing.T) {
	inputURL := "http://example.com"
	defaultPort := false

	expectedURL := "example.com"

	result, err := SanitizeURL(inputURL, defaultPort)

	assert.Equal(t, expectedURL, result)
	assert.Nil(t, err)
}
