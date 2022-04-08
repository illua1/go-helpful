package helpful

import(
  "golang.org/x/exp/constraints"
)

type Values interface {
	constraints.Float | constraints.Integer | byte
}

func Cast[A, B Values](a []A) (b []B) {
	b = make([]B, len(a))
	for i := range a {
		b[i] = B(a[i])
	}
	return
}


func Fill[T any](in []T, method func(index int) T) {
	for i := range in {
		in[i] = method(i)
	}
}

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