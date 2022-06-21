package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
)

func Abc[Value value.Values](a Value) Value {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func Lerp[T, F value.Values](a, b T, f F) T {
	return a + T(F(b-a)*f)
}

func UnLerp[T, F value.Values](a, b T, x T) F {
	return F((x - a) / (b - a))
}
