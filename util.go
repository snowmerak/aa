package aa

func Box[T any](a T) *T {
	return &a
}
