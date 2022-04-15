package helpful_vector_matrix

import (
	"fmt"
	value "github.com/illua1/go-helpful"
	"image"
)

type matrixSlise[Value value.Values] struct {
	src   Matrix_Main_Functions[Value]
	bound image.Rectangle
}

func (matrix_slise *matrixSlise[Value]) String() (as string) {
	as = "["
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		as += "Row[" + fmt.Sprint(x) + "]:{"
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			as += fmt.Sprint(matrix_slise.src.Get(x, y)) + " "
		}
		as = as[:len(as)-1]
		as += "}"
	}
	as += "]"
	return
}

func (matrix_slise *matrixSlise[Value]) Bound() image.Rectangle {
	return matrix_slise.bound
}

func (matrix_slise *matrixSlise[Value]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return &matrixSlise[Value]{
		src: matrix_slise.src,
		bound: matrix_slise.bound.Intersect(
			image.Rect(
				min_x,
				min_y,
				max_x,
				max_y,
			).Add(
				matrix_slise.bound.Min,
			),
		),
	}
}

func (matrix_slise *matrixSlise[Value]) Get(x, y int) Value {
	return matrix_slise.src.Get(x+matrix_slise.bound.Min.X, y+matrix_slise.bound.Min.Y)
}

func (matrix_slise *matrixSlise[Value]) Set(x, y int, value Value) {
	matrix_slise.src.Set(x+matrix_slise.bound.Min.X, y+matrix_slise.bound.Min.Y, value)
}

func (matrix_slise *matrixSlise[Value]) Size() (x, y int) {
	return matrix_slise.bound.Dx(), matrix_slise.bound.Dy()
}

func (matrix_slise *matrixSlise[Value]) Minor(x_, y_ int) Matrix_Main_Functions[Value] {
	size_x, size_y := matrix_slise.Size()
	var ret MatrixMutable[Value] = matrix_slise.Slise(0, 0, size_x-1, size_y-1).Mutable()
	if (0 <= x_) && (x_ < size_x) {
		if (0 <= y_) && (y_ < size_y) {
			for x := 0; x < x_; x++ {
				for y := 0; y < y_; y++ {
					ret.Set(x, y, matrix_slise.Get(x, y))
				}
				for y := y_ + 1; y < size_x; y++ {
					ret.Set(x, y-1, matrix_slise.Get(x, y))
				}
			}
			for x := x_ + 1; x < size_x; x++ {
				for y := 0; y < y_; y++ {
					ret.Set(x-1, y, matrix_slise.Get(x, y))
				}
				for y := y_ + 1; y < size_y; y++ {
					ret.Set(x-1, y-1, matrix_slise.Get(x, y))
				}
			}
		}
	}
	return &ret
}

func (matrix_slise *matrixSlise[Value]) Body() []Value {
	var ret = make([]Value, matrix_slise.bound.Dx()*matrix_slise.bound.Dy())
	var pointer int
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			ret[pointer] = matrix_slise.Get(x, y)
			pointer++
		}
	}
	return ret
}

func (matrix_slise *matrixSlise[Value]) Fill(value Value) {
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			matrix_slise.src.Set(x, y, value)
		}
	}
}

func (matrix_slise *matrixSlise[Value]) Scale(size Value) {
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			matrix_slise.src.Set(x, y, matrix_slise.src.Get(x, y)*size)
		}
	}
}

func (matrix_slise *matrixSlise[Value]) Mutable() MatrixMutable[Value] {
	var m = MatrixMutableNew[Value](matrix_slise.Size())
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		var pointer_x int
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			var pointer_y int
			m.Set(pointer_x, pointer_y, matrix_slise.Get(x, y))
			pointer_y++
		}
		pointer_x++
	}
	return m
}

func (matrix_slise *matrixSlise[Value]) FillAs(values func(x, y int) Value) {
	x_, y_ := matrix_slise.Size()
	for x := 0; x < x_; x++ {
		for y := 0; y < y_; y++ {
			matrix_slise.Set(x, y, values(x, y))
		}
	}
	return
}

func (matrix_slise *matrixSlise[Value]) FillTo(in func(x, y int, value Value)) {
	x_, y_ := matrix_slise.Size()
	for x := 0; x < x_; x++ {
		for y := 0; y < y_; y++ {
			in(x, y, matrix_slise.Get(x, y))
		}
	}
	return
}
