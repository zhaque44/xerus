package collections

import "errors"

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
