package collections

import (
	"errors"
	"reflect"
)

func Map[T, U any](s []T, f func(T) (U, error)) ([]U, error) {
	r := make([]U, len(s))
	for i, v := range s {
		u, err := f(v)
		if err != nil {
			return nil, err
		}
		r[i] = u
	}
	return r, nil
}

func Values(m map[int]string) ([]string, error) {
	if len(m) == 0 {
		return nil, errors.New("map is empty")
	}

	values := make([]string, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values, nil
}

func ContainsKey(m interface{}, key interface{}) bool {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return false
	}

	if v.MapIndex(reflect.ValueOf(key)).IsValid() {
		return true
	}

	return false
}

func ContainsValue(m interface{}, value interface{}) bool {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return false
	}

	for _, k := range v.MapKeys() {
		if v.MapIndex(k).Interface() == value {
			return true
		}
	}

	return false
}

func foldLeft(slice1, slice2 []int, init int, op func(int, int, int) int) int {
	result := init
	for i := 0; i < len(slice1) && i < len(slice2); i++ {
		result = op(result, slice1[i], slice2[i])
	}
	return result
}
