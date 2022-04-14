package helpful_compare

func Compare[T comparable](a, b []T) bool {
	if (a == nil) != (b == nil) {
		return false
	} else {
		if (a == nil) || (b == nil) {
			return false
		}
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
