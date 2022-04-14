package helpful_sort

import (
	"github.com/illua1/go-helpful"
)

func ClampF[T helpful.Values](min, max, value T) T {
	if (min <= value) && (value <= max) {
		return value
	}
	if min > value {
		return min
	}
	return max
}

func Clamp[T any](min, max, value T, condition func(a, b T) bool) T {
	if condition(min, value) && condition(max, value) {
		return value
	}
	if condition(min, value) {
		return min
	}
	return max
}
