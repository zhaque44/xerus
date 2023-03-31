package machinelearning

import (
	"errors"
)

func intersection[T comparable](slice1, slice2 []T) ([]T, error) {
	if slice1 == nil || slice2 == nil {
		return nil, errors.New("input slices cannot be nil")
	}

	set := make(map[T]bool)
	var result []T

	for _, v := range slice1 {
		set[v] = true
	}

	for _, v := range slice2 {
		if set[v] {
			result = append(result, v)
		}
	}

	return result, nil
}
