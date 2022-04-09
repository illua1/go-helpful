package helpful_slise

import(
  "github.com/illua1/go-helpful"
)

// Cast will do type cast each elements of input slise to second type.
func Cast[A, B helpful.Values](a []A) (b []B) {
	b = make([]B, len(a))
	for i := range a {
		b[i] = B(a[i])
	}
	return
}

// Fill will do rewrite input slise by function output.
func Fill[T any](in []T, method func(index int) T) {
	for i := range in {
		in[i] = method(i)
	}
}

// Join return slise that will contained all element of input sliset, order will saved.
func Join[T any](slises ...[]T) (out []T) {
	var length int = 0
	for i := range slises {
		length += len(slises[i])
	}
	out = make([]T, length)
	length = 0
	for i := range slises {
		copy(
			out[length:length+len(slises[i])],
			slises[i],
		)
		length += len(slises[i])
	}
	return
}