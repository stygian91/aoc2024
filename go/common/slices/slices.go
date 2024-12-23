package slices

func LastIndexFunc[T any](s []T, fn func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return i
		}
	}

	return -1
}
