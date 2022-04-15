package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
)

func MatrixMutableNew[Value value.Values](size_x, size_y int) MatrixMutable[Value] {
	return MatrixMutable[Value]{
		a:      make([]Value, size_x*size_y),
		size_x: uint(size_x),
		size_y: uint(size_y),
	}
}
