package collections

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
