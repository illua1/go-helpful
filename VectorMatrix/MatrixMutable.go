package helpful_vector_matrix

import (
	"fmt"
	value "github.com/illua1/go-helpful"
	slise "github.com/illua1/go-helpful/Slise"
	"image"
)

/*

  Matrix mutable {
    a = X Y Z X Y Z X Y Z
    size_x = 3
    size_y = 3
  }

*/

type MatrixMutable[
	Value value.Values,
] struct {
	a              []Value
	size_x, size_y uint
}

func (matrix_mutable *MatrixMutable[Value]) String() (as string) {
	as = "["
	x_, y_ := matrix_mutable.Size()
	for x := 0; x < x_; x++ {
		as += "Row[" + fmt.Sprint(x) + "]:{"
		for y := 0; y < y_; y++ {
			as += fmt.Sprint(matrix_mutable.Get(x, y)) + " "
		}
		as = as[:len(as)-1]
		as += "}"
	}
	as += "]"
	return
}

func (matrix_mutable *MatrixMutable[Value]) Get(x, y int) Value {
	return matrix_mutable.a[x*int(matrix_mutable.size_x)+y]
}

func (matrix_mutable *MatrixMutable[Value]) Set(x, y int, value Value) {
	matrix_mutable.a[x*int(matrix_mutable.size_x)+y] = value
}

func (matrix_mutable *MatrixMutable[Value]) Size() (x, y int) {
	return int(matrix_mutable.size_x), int(matrix_mutable.size_y)
}

func (matrix_mutable *MatrixMutable[Value]) Bound() image.Rectangle {
	return image.Rect(0, 0, int(matrix_mutable.size_x), int(matrix_mutable.size_y))
}

func (matrix_mutable *MatrixMutable[Value]) Fill(value Value) {
	slise.CopyTo(matrix_mutable.a, value)
}

func (matrix_mutable *MatrixMutable[Value]) Scale(size Value) {
	for i := range matrix_mutable.a {
		matrix_mutable.a[i] *= size
	}
}

func (matrix_mutable *MatrixMutable[Value]) Minor(x_, y_ int) Matrix_Main_Functions[Value] {
	var ret MatrixMutable[Value]
	size_x, size_y := matrix_mutable.Size()
	if (0 <= x_) && (x_ < size_x) {
		if (0 <= y_) && (y_ < size_y) {
			ret.size_x = uint(size_x - 1)
			ret.size_y = uint(size_y - 1)
			ret.a = make([]Value, (size_x-1)*(size_y-1))
			for x := 0; x < x_; x++ {
				for y := 0; y < y_; y++ {
					ret.Set(x, y, matrix_mutable.Get(x, y))
				}
				for y := y_ + 1; y < size_x; y++ {
					ret.Set(x, y-1, matrix_mutable.Get(x, y))
				}
			}
			for x := x_ + 1; x < size_x; x++ {
				for y := 0; y < y_; y++ {
					ret.Set(x-1, y, matrix_mutable.Get(x, y))
				}
				for y := y_ + 1; y < size_y; y++ {
					ret.Set(x-1, y-1, matrix_mutable.Get(x, y))
				}
			}
		}
	}
	return &ret
}

func (matrix_mutable *MatrixMutable[Value]) Body() (ret []Value) {
	ret = make([]Value, len(matrix_mutable.a))
	copy(ret, matrix_mutable.a)
	return
}

func (matrix_mutable *MatrixMutable[Value]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return &matrixSlise[Value]{
		src:   matrix_mutable,
		bound: matrix_mutable.Bound().Intersect(image.Rect(min_x, min_y, max_x, max_y)),
	}
}

func (matrix_mutable *MatrixMutable[Value]) Mutable() MatrixMutable[Value] {
	var a []Value = make([]Value, len(matrix_mutable.a))
	copy(a, matrix_mutable.a)
	return MatrixMutable[Value]{
		a:      a,
		size_x: matrix_mutable.size_x,
		size_y: matrix_mutable.size_y,
	}
}

func (matrix_mutable *MatrixMutable[Value]) FillAs(values func(x, y int) Value) {
	for x := 0; x < int(matrix_mutable.size_x); x++ {
		for y := 0; y < int(matrix_mutable.size_y); y++ {
			matrix_mutable.Set(x, y, values(x, y))
		}
	}
	return
}

func (matrix_mutable *MatrixMutable[Value]) FillTo(in func(x, y int, value Value)) {
	for x := 0; x < int(matrix_mutable.size_x); x++ {
		for y := 0; y < int(matrix_mutable.size_y); y++ {
			in(x, y, matrix_mutable.Get(x, y))
		}
	}
	return
}
