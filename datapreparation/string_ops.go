package datapreparation

import (
	"strings"

	"github.com/pkg/errors"
)

func ContainsPrefix(input string, prefixes ...string) bool {
	if input == "" {
		return false
	}

	inputLower := strings.ToLower(input)
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToLower(prefix), inputLower) {
			return true
		}
	}

	return false
}

func SplitAndTrimWhitespace(inputString, splitchar string) ([]string, error) {
	if splitchar == "" {
		return nil, errors.New("splitchar cannot be empty")
	}

	splitStrings := strings.Split(inputString, splitchar)
	var result []string

	for _, trimspace := range splitStrings {
		result = append(result, strings.TrimSpace(trimspace))
	}

	return result, nil
}
