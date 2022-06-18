package helpful_vector_matrix

import (
	"fmt"
	value "github.com/illua1/go-helpful"
	"image"
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

type Matrix[
	Value value.Values,
	Row value.Arrays[Value],
	Column value.Arrays[Row],
] struct {
	A Column
}

func (matrix *Matrix[Value, Row, Column]) String() (as string) {
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

func (matrix *Matrix[Value, Row, Column]) Body() (ret []Value) {
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

func (matrix *Matrix[Value, Row, Column]) Bound() image.Rectangle {
	return image.Rect(0, 0, len(matrix.A), len(matrix.A[0]))
}

func (matrix *Matrix[Value, Row, Column]) Get(x, y int) Value {
	return matrix.A[x][y]
}

func (matrix *Matrix[Value, Row, Column]) Set(x, y int, value Value) {
	matrix.A[x][y] = value
}

func (matrix *Matrix[Value, Row, Column]) Size() (x, y int) {
	return len(matrix.A), len(matrix.A[0])
}

func (matrix *Matrix[Value, Row, Column]) Fill(value Value) {
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			matrix.A[x][y] = value
		}
	}
}

func (matrix *Matrix[Value, Row, Column]) Scale(size Value) {
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			matrix.A[x][y] *= size
		}
	}
}

/*

  Matrix3x3 {
    1   2   3
    4   5   6
    7   8   9
  }

  -> 0, 0
  -> {
    0   0   0
    0   5   6
    0   8   9
  } -> {
    5   6
    8   9
  }

  -> 1, 1
  -> {
    1   0   3
    0   0   0
    7   0   9
  } -> {
    1   3
    7   9
  }

  -> 2, 0
  -> {
    0   0   0
    4   5   0
    7   8   0
  } -> {
    4   5
    7   8
  }

*/

func (matrix *Matrix[Value, Row, Column]) Minor(x_, y_ int) Matrix_Main_Functions[Value] {
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
	return &ret
}

func (matrix *Matrix[Value, Row, Column]) Slise(min_x, min_y, max_x, max_y int) Matrix_Main_Functions[Value] {
	return Matrix_Main_Functions[Value](&matrixSlise[Value]{
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

func (matrix *Matrix[Value, Row, Column]) Mutable() (ret MatrixMutable[Value]) {
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

func (matrix *Matrix[Value, Row, Column]) FillAs(values func(x, y int) Value) {
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			matrix.A[x][y] = values(x, y)
		}
	}
	return
}

func (matrix *Matrix[Value, Row, Column]) FillTo(in func(x, y int, value Value)) {
	for x := 0; x < len(matrix.A); x++ {
		for y := 0; y < len(matrix.A[0]); y++ {
			in(x, y, matrix.A[x][y])
		}
	}
	return
}

func (matrix *Matrix[Value, Row, Column]) MulVector(a Vector[Value, Row]) Vector[Value, Row] {
	if len(matrix.A) != len(matrix.A[0]) {
		return Vector[Value, Row]{}
	}
	var ret Vector[Value, Row]
	for i := 0; i < len(matrix.A); i++ {
		for x := 0; x < len(matrix.A[0]); x++ {
			ret.A[i] += a.A[x] * matrix.A[i][x]
		}
	}
	return ret
}

func (matrix *Matrix[Value, Row, Column]) Determinant() Value {
	if len(matrix.A) != len(matrix.A[0]) {
		return 0
	}
	if len(matrix.A) > 1 {
		var mutable = matrix.Mutable()
		return (&mutable).Determinant()
	}
	return matrix.A[0][0]
}

func (matrix Matrix[Value, Row, Column]) Invert() Matrix[Value, Row, Column] {
	if len(matrix.A) != len(matrix.A[0]) {
		return Matrix[Value, Row, Column]{}
	}
	var ret Matrix[Value, Row, Column]
	var mutable = matrix.Mutable()
	var determinamnt = matrix.Determinant()
	if determinamnt == 0 {
		return Matrix[Value, Row, Column]{}
	}
	for x := 0; x < len(ret.A); x++ {
		for y := 0; y < len(ret.A[0]); y++ {
			ret.A[y][x] = (Value(1-((x+y)%2)*2) * mutable.Minor(x, y).Determinant()) / determinamnt
		}
	}
	return ret
}
