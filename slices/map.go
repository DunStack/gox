package slices

func Map[S ~[]E, E any, T any](s S, fn func(e E) T) []T {
	var values []T
	for _, e := range s {
		values = append(values, fn(e))
	}
	return values
}
