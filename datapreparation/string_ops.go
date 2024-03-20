package datapreparation

import (
	"encoding/json"
	"fmt"
	"net/url"
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

func converter(data interface{}, urlStr string) (interface{}, error) {
	switch data.(type) {
	case string:
		if urlStr != "" {
			parts := []string{urlStr}
			u, err := url.Parse(urlStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing URL: %v", err)
			}
			parts = append(parts, u.Path[1:])
			return parts, nil
		}
		var parsedData interface{}
		if err := json.Unmarshal([]byte(data.(string)), &parsedData); err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON data: %v", err)
		}
		return parsedData, nil
	default:
		var parts []string
		if urlStr != "" {
			parts = append(parts, urlStr)
		}
		for _, part := range data.(map[string]interface{}) {
			parts = append(parts, fmt.Sprintf("%v", part))
		}
		return parts, nil
	}
}
