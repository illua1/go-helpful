package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
	"math"
)

/*
  Rotate matrix (Auler) (X Y Z)
*/

func Rotate3x3_XYZ[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_x[Value](angle[0]),
	).Mull(
		Rotate3x3_y[Value](angle[1]),
	).Mull(
		Rotate3x3_z[Value](angle[2]),
	)
}

func Rotate3x3_XZY[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_x[Value](angle[0]),
	).Mull(
		Rotate3x3_z[Value](angle[2]),
	).Mull(
		Rotate3x3_y[Value](angle[1]),
	)
}

func Rotate3x3_YZX[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_y[Value](angle[1]),
	).Mull(
		Rotate3x3_z[Value](angle[2]),
	).Mull(
		Rotate3x3_x[Value](angle[0]),
	)
}

func Rotate3x3_YXZ[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_y[Value](angle[1]),
	).Mull(
		Rotate3x3_x[Value](angle[0]),
	).Mull(
		Rotate3x3_z[Value](angle[2]),
	)
}

func Rotate3x3_ZYX[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_z[Value](angle[2]),
	).Mull(
		Rotate3x3_y[Value](angle[1]),
	).Mull(
		Rotate3x3_x[Value](angle[0]),
	)
}

func Rotate3x3_ZXY[Value value.Values, Angle value.Values](angle [3]Angle) Matrix[Value, [3]Value, [3][3]Value] {
	return Matrix3x3[Value]().Mull(
		Rotate3x3_z[Value](angle[2]),
	).Mull(
		Rotate3x3_x[Value](angle[0]),
	).Mull(
		Rotate3x3_y[Value](angle[1]),
	)
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
