package helpful_vector_matrix

import (
	"fmt"
	value "github.com/illua1/go-helpful"
	"image"
)

type Matrix_Main_Functions[Value value.Values] interface {
	String() string

	Get(x, y int) Value
	Set(s, y int, value Value)

	Bound() image.Rectangle
	Size() (x, y int)
	Body() []Value

	Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value]
	Mutable() MatrixMutable[Value]

	Fill(value Value)
	FillTo(in func(x, y int, value Value))
	FillAs(values func(x, y int) Value)

	Scale(size Value)

	Minor(x, y int) Matrix_Main_Functions[Value]
}

func init() {
	var a Matrix_Main_Functions[int]
	{
		a_l := Matrix2x2[int]()
		a = &a_l
	}
	{
		a_f := Matrix2x2[int]()
		a = a_f.Slise(0, 0, 2, 2)
	}
	{
		a_f := Matrix2x2[int]()
		a_l := a_f.Mutable()
		a = &a_l
	}

	if false {
		fmt.Println(a)
	}
}

type MatrixOperations[Value value.Values, Row value.Arrays[Value], Column value.Arrays[Row]] interface {
	Matrix[Value, Row, Column] | matrixSlise[Value] | MatrixMutable[Value]
}

func MatrixEqualSize[Value_a value.Values, Value_b value.Values](
	matrix_a Matrix_Main_Functions[Value_a],
	matrix_b Matrix_Main_Functions[Value_b],
) bool {
	var x_1, y_1 = matrix_a.Size()
	var x_2, y_2 = matrix_a.Size()
	return (x_1 == x_2) && (y_1 == y_2)
}

func MatrixIsEqual[Value_a value.Values, Value_b value.Values](
	matrix_a Matrix_Main_Functions[Value_a],
	matrix_b Matrix_Main_Functions[Value_b],
) bool {
	if !MatrixEqualSize(matrix_a, matrix_b) {
		return false
	}
	x_, y_ := matrix_a.Size()
	for x := 0; x < x_; x++ {
		for y := 0; y < y_; y++ {
			if matrix_a.Get(x, y) != Value_a(matrix_b.Get(x, y)) {
				return false
			}
		}
	}
	return true
}

func MatrixWrite[Value_a value.Values, Value_b value.Values](
	matrix_a Matrix_Main_Functions[Value_a],
	matrix_b Matrix_Main_Functions[Value_b],
) {
	if !MatrixEqualSize(matrix_a, matrix_b) {
		return
	}
	x_, y_ := matrix_a.Size()
	for x := 0; x < x_; x++ {
		for y := 0; y < y_; y++ {
			matrix_a.Set(x, y, Value_a(matrix_b.Get(x, y)))
		}
	}
}
