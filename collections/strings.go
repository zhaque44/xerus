package collections

import (
	"fmt"
	"strconv"
	"strings"
)

// IntSliceFromString converts a comma-separated string of integers into a slice of integers.
func IntSliceFromString(s string) ([]int, error) {
	if s == "" {
		return nil, nil
	}

	parts := strings.Split(s, ",")
	result := make([]int, len(parts))
	validIntFound := false

	for i, v := range parts {
		vTrim := strings.TrimSpace(v)
		num, err := strconv.Atoi(vTrim)
		if err != nil {
			return nil, err // Return an error if any conversion fails
		}
		result[i] = num
		validIntFound = true
	}

	if !validIntFound {
		return nil, fmt.Errorf("no valid integers found in the input string")
	}

	return result, nil
}
