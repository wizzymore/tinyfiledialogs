package tinyfiledialogs

func tern[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
