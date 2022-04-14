package helpful_vector_matrix

import (
	"fmt"
	value "github.com/illua1/go-helpful"
	slise "github.com/illua1/go-helpful/Slise"
	"image"
	"math"
)

/*

  Matrix3x_3{
    X X X
    Y Y Y
    Z Z Z
  }

  Row{
    X
    Y
    Z
  }

  Column{
    Row
    Row
    Row
  }

  Matrix[0] -> Row[0] -> X

  Matrix[X][Y]

*/

type Matrix_Main_Functions[Value value.Values] interface {
	Get(x, y int) Value
	Set(s, y int, value Value)

	Bound() image.Rectangle
	Size() (x, y int)
	Body() []Value

	Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value]
	Fill(value Value)

	Minor(x, y int) Matrix_Main_Functions[Value]
}

func init() {
	var a Matrix_Main_Functions[int]
	a = Matrix2x2[int]()
	a = Matrix2x2[int]().Slise(0, 0, 2, 2)
	a = Matrix2x2[int]().Mutable()

	if a != nil {
		fmt.Println(a)
	}
}

type Matrix[
	Value value.Values,
	Row value.Arrays[Value],
	Column value.Arrays[Row],
] struct {
	A Column
}

func (matrix Matrix[Value, Row, Column]) String() (as string) {
	as = "["
	for x := 0; x < len(matrix.A); x++ {
		as += "Row[" + fmt.Sprint(x) + "]:{"
		for y := 0; y < len(matrix.A[0]); y++ {
			as += fmt.Sprint(matrix.A[x][y]) + " "
		}
		as = as[:len(as)-1]
		as += "}"
	}
	as += "]"
	return
}

func Matrix2x2[Value value.Values]() Matrix[Value, [2]Value, [2][2]Value] {
	var m Matrix[Value, [2]Value, [2][2]Value]
	m.A[0][0] = 1
	m.A[1][1] = 1
	return m
}

func Matrix3x3[Value value.Values]() Matrix[Value, [3]Value, [3][3]Value] {
	var m Matrix[Value, [3]Value, [3][3]Value]
	m.A[0][0] = 1
	m.A[1][1] = 1
	m.A[2][2] = 1
	return m
}

func Matrix4x4[Value value.Values]() Matrix[Value, [4]Value, [4][4]Value] {
	var m Matrix[Value, [4]Value, [4][4]Value]
	m.A[0][0] = 1
	m.A[1][1] = 1
	m.A[2][2] = 1
	m.A[3][3] = 1
	return m
}

func Matrix5x5[Value value.Values]() Matrix[Value, [5]Value, [5][5]Value] {
	var m Matrix[Value, [5]Value, [5][5]Value]
	m.A[0][0] = 1
	m.A[1][1] = 1
	m.A[2][2] = 1
	m.A[3][3] = 1
	m.A[4][4] = 1
	return m
}

func Matrix6x6[Value value.Values]() Matrix[Value, [6]Value, [6][6]Value] {
	var m Matrix[Value, [6]Value, [6][6]Value]
	m.A[0][0] = 1
	m.A[1][1] = 1
	m.A[2][2] = 1
	m.A[3][3] = 1
	m.A[4][4] = 1
	m.A[5][5] = 1
	return m
}

/*

  Matrix3x3{
    X   X   X
    Y   Y   Y
    Z   Z   Z
  }

  Mull to

  Matrix3x3{
    A   A   A
    B   B   B
    C   C   C
  }

  ->

  Matrix3x3{
    XA+XB+XC   XA+XB+XC   XA+XB+XC
    YA+YB+YC   YA+YB+YC   YA+YB+YC
    ZA+ZB+ZC   ZA+ZB+ZC   ZA+ZB+ZC
  }

  Row[0,1,2]{
    XA+XB+XC
    YA+YB+YC
    ZA+ZB+ZC
  }

*/

func (matrix Matrix[Value, Row, Column]) Mull(B Matrix[Value, Row, Column]) (C Matrix[Value, Row, Column]) {
	var len_ int
	if len(matrix.A) > len(matrix.A[0]) {
		len_ = len(matrix.A[0])
	} else {
		len_ = len(matrix.A)
	}
	for x := 0; x < len_; x++ {
		for y := 0; y < len_; y++ {
			for i := 0; i < len_; i++ {
				C.A[x][y] += matrix.A[i][x] * B.A[y][i]
			}
		}
	}
	return
}

func Rotate3x3[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	var m = Matrix3x3[Value]()

	m = m.Mull(Rotate3x3_x[Value](angle[0]))
	m = m.Mull(Rotate3x3_y[Value](angle[1]))
	m = m.Mull(Rotate3x3_z[Value](angle[2]))

	return m
}

/*

  Matrix3x3 Rotate : X [Y and Z] {
    X :  1    0    0
    Y :  0   cos -sin
    Z :  0   sin  cos
  }

  Row[0]{
     1
     0
     0
  }

  Row[1]{
     0
    cos
    sin
  }

  Row[2]{
     0
   -sin
    cos
  }

*/

func Rotate3x3_x[Value value.Values, Angle value.Values](angle Angle) Matrix[Value, [3]Value, [3][3]Value] {
	var m Matrix[Value, [3]Value, [3][3]Value]
	m.A[0][0] = 1
	m.A[1][1] = Value(math.Cos(float64(angle)))
	m.A[1][2] = Value(-math.Sin(float64(angle)))
	m.A[2][1] = Value(math.Sin(float64(angle)))
	m.A[2][2] = Value(math.Cos(float64(angle)))
	return m
}

/*

  Matrix3x3 Rotate : Y [X and Z] {
    X : cos   0   sin
    Y :  0    1    0
    Z :-sin   0   cos
  }

  Row[0]{
    cos
     0
   -sin
  }

  Row[1]{
     0
     1
     0
  }

  Row[2]{
    sin
     0
    cos
  }

*/

func Rotate3x3_y[Value value.Values, Angle value.Values](angle Angle) Matrix[Value, [3]Value, [3][3]Value] {
	var m Matrix[Value, [3]Value, [3][3]Value]
	m.A[0][0] = Value(math.Cos(float64(angle)))
	m.A[0][2] = Value(-math.Sin(float64(angle)))
	m.A[1][1] = 1
	m.A[2][0] = Value(math.Sin(float64(angle)))
	m.A[2][2] = Value(math.Cos(float64(angle)))
	return m
}

/*

  Matrix3x3 Rotate : Z [X and Y] {
    X : cos -sin   0
    Y : sin  cos   0
    Z :  0    0    1
  }

  Row[0]{
    cos
    sin
     0
  }

  Row[1]{
   -sin
    cos
     0
  }

  Row[2]{
     0
     0
     1
  }

*/

func Rotate3x3_z[Value value.Values, Angle value.Values](angle Angle) Matrix[Value, [3]Value, [3][3]Value] {
	var m Matrix[Value, [3]Value, [3][3]Value]
	m.A[0][0] = Value(math.Cos(float64(angle)))
	m.A[0][1] = Value(math.Sin(float64(angle)))
	m.A[1][0] = Value(-math.Sin(float64(angle)))
	m.A[1][1] = Value(math.Cos(float64(angle)))
	m.A[2][2] = 1
	return m
}

func (matrix *Matrix[Value, Row, Column]) Transposed() (C Matrix[Value, Row, Column]) {
	if len(matrix.A) != len(matrix.A[0]) {
		return
	}
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			C.A[x][y] = matrix.A[y][x]
		}
	}
	return
}

func (matrix Matrix[Value, Row, Column]) Body() (ret []Value) {
	ret = make([]Value, len(matrix.A)*len(matrix.A[0]))
	var pointer int
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			ret[pointer] = matrix.A[x][y]
			pointer++
		}
	}
	return
}

func (matrix Matrix[Value, Row, Column]) Bound() image.Rectangle {
	return image.Rect(0, 0, len(matrix.A), len(matrix.A[0]))
}

func (matrix Matrix[Value, Row, Column]) Get(x, y int) Value {
	return matrix.A[x][y]
}

func (matrix Matrix[Value, Row, Column]) Set(x, y int, value Value) {
	matrix.A[x][y] = value
}

func (matrix Matrix[Value, Row, Column]) Size() (x, y int) {
	return len(matrix.A), len(matrix.A[0])
}

func (matrix Matrix[Value, Row, Column]) Fill(value Value) {
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			matrix.A[x][y] = value
		}
	}
}

/*

  Matrix3x3 {
    1   2   3
    4   5   6
    7   8   9
  }

  -> 0, 0 -> {
    0   0   0
    0   5   6
    0   8   9
  } -> {
    5   6
    8   9
  }

  -> 1, 1 -> {
    1   0   3
    0   0   0
    7   0   9
  } -> {
    1   3
    7   9
  }

  -> 2, 0 -> {
    0   0   0
    4   5   0
    7   8   0
  } -> {
    4   5
    7   8
  }

*/

func (matrix Matrix[Value, Row, Column]) Minor(x_, y_ int) Matrix_Main_Functions[Value] {
	var ret Matrix[Value, Row, Column]
	if (0 <= x_) && (x_ < len(matrix.A)) {
		if (0 <= y_) && (y_ < len(matrix.A[0])) {
			for x := 0; x < x_; x++ {
				for y := 0; y < y_; y++ {
					ret.A[x][y] = matrix.A[x][y]
				}
				for y := y_ + 1; y < len(matrix.A[0]); y++ {
					ret.A[x][y-1] = matrix.A[x][y]
				}
			}
			for x := x_ + 1; x < len(matrix.A); x++ {
				for y := 0; y < y_; y++ {
					ret.A[x-1][y] = matrix.A[x][y]
				}
				for y := y_ + 1; y < len(matrix.A[0]); y++ {
					ret.A[x-1][y-1] = matrix.A[x][y]
				}
			}
		}
	}
	return ret
}

/*
func(matrix *Matrix[Value, Row, Column])Determinant()Value{
  if len(matrix.A) != len(matrix.A[0]) {
    var v Value
    return v
  }
  if len(matrix.A) == 2 {
    var i = 1
    return matrix.A[0][0] * matrix.A[i][i] - matrix.A[0][i] * matrix.A[i][0]
  }else{
    //var buffer Matrix[Value, Row, Column]
    //for diagonal  := 0; diagonal < len(matrix.A)-2; diagonal++ {
      //for x_main := 0; x_main < len(matrix.A); x_main++ {
        //buffer = matrix.Minor(x_main, 0)

      //}
    //}
  }
}
*/

type matrixSlise[Value value.Values] struct {
	src   Matrix_Main_Functions[Value]
	bound image.Rectangle
}

func (matrix Matrix[Value, Row, Column]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return Matrix_Main_Functions[Value](matrixSlise[Value]{
		src: Matrix_Main_Functions[Value](matrix),
		bound: image.Rect(
			0,
			0,
			len(matrix.A),
			len(matrix.A[0]),
		).Intersect(
			image.Rect(min_x, min_y, max_x, max_y),
		),
	})
}

func (matrix Matrix[Value, Row, Column]) Mutable() (ret MatrixMutable[Value]) {
	ret.size_x = uint(len(matrix.A))
	ret.size_y = uint(len(matrix.A[0]))
	ret.a = make([]Value, ret.size_x*ret.size_y)
	var pointer int
	for x := 0; x < int(ret.size_x); x++ {
		for y := 0; y < int(ret.size_y); y++ {
			ret.a[pointer] = matrix.A[x][y]
			pointer++
		}
	}
	return
}

func (matrix_slise matrixSlise[Value]) String() (as string) {
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

func (matrix_slise matrixSlise[Value]) Bound() image.Rectangle {
	return matrix_slise.bound
}

func (matrix_slise matrixSlise[Value]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return matrixSlise[Value]{
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

func (matrix_slise matrixSlise[Value]) Get(x, y int) Value {
	return matrix_slise.src.Get(x+matrix_slise.bound.Min.X, y+matrix_slise.bound.Min.Y)
}

func (matrix_slise matrixSlise[Value]) Set(x, y int, value Value) {
	matrix_slise.src.Set(x+matrix_slise.bound.Min.X, y+matrix_slise.bound.Min.Y, value)
}

func (matrix_slise matrixSlise[Value]) Size() (x, y int) {
	return matrix_slise.bound.Dx(), matrix_slise.bound.Dy()
}

func (matrix_slise matrixSlise[Value]) Minor(x, y int) Matrix_Main_Functions[Value] {
	return matrix_slise.src.Minor(x, y)
}

func (matrix_slise matrixSlise[Value]) Body() []Value {
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

func (matrix_slise matrixSlise[Value]) Fill(value Value) {
	for x := matrix_slise.bound.Min.X; x < matrix_slise.bound.Max.X; x++ {
		for y := matrix_slise.bound.Min.Y; y < matrix_slise.bound.Max.Y; y++ {
			matrix_slise.src.Set(x, y, value)
		}
	}
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

func (matrix_mutable MatrixMutable[Value]) Get(x, y int) Value {
	return matrix_mutable.a[x*int(matrix_mutable.size_x)+y]
}

func (matrix_mutable MatrixMutable[Value]) Set(x, y int, value Value) {
	matrix_mutable.a[x*int(matrix_mutable.size_x)+y] = value
}

func (matrix_mutable MatrixMutable[Value]) Size() (x, y int) {
	return int(matrix_mutable.size_x), int(matrix_mutable.size_y)
}

func (matrix_mutable MatrixMutable[Value]) Bound() image.Rectangle {
	return image.Rect(0, 0, int(matrix_mutable.size_x), int(matrix_mutable.size_y))
}

func (matrix_mutable MatrixMutable[Value]) Fill(value Value) {
	slise.CopyTo(matrix_mutable.a, value)
}

func (matrix_mutable MatrixMutable[Value]) Minor(x_, y_ int) Matrix_Main_Functions[Value] {
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
	return ret
}

func (matrix_mutable MatrixMutable[Value]) Body() (ret []Value) {
	ret = make([]Value, len(matrix_mutable.a))
	copy(ret, matrix_mutable.a)
	return
}

func (matrix_mutable MatrixMutable[Value]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return matrixSlise[Value]{
		src:   matrix_mutable,
		bound: matrix_mutable.Bound().Intersect(image.Rect(min_x, min_y, max_x, max_y)),
	}
}
