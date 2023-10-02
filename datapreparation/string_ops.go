package datapreparation

import "strings"

func ContainsPrefix(input string, prefixes ...string) bool {
	inputLower := strings.ToLower(input)
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToLower(prefix), inputLower) {
			return true
		}
	}

	return false
}
