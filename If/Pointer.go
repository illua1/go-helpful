package helpful_ternary

// Simple Pointer operator.
func Pointer[T any](in T) *T {
	return &in
}
