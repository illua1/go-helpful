package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
)

type Vector[
	Value value.Values,
	Length value.Arrays[Value],
] struct {
	A Length
}

func Vector3[Value value.Values](fill Value) Vector[Value, [3]Value] {
	var v Vector[Value, [3]Value]
	v.A[0] = fill
	v.A[1] = fill
	v.A[2] = fill
	return v
}

func Vector2[Value value.Values](fill Value) Vector[Value, [2]Value] {
	var v Vector[Value, [2]Value]
	v.A[0] = fill
	v.A[1] = fill
	return v
}

func (vector *Vector[Value, Length]) Add(a Vector[Value, Length]) {
	for i := 0; i < len(vector.A); i++ {
		vector.A[i] = vector.A[i] + a.A[i]
	}
}

func (vector *Vector[Value, Length]) Sub(a Vector[Value, Length]) {
	for i := 0; i < len(vector.A); i++ {
		vector.A[i] = vector.A[i] - a.A[i]
	}
}

func (vector *Vector[Value, Length]) Mull(a Vector[Value, Length]) {
	for i := 0; i < len(vector.A); i++ {
		vector.A[i] = vector.A[i] * a.A[i]
	}
}

func (vector *Vector[Value, Length]) Div(a Vector[Value, Length]) {
	for i := 0; i < len(vector.A); i++ {
		vector.A[i] = vector.A[i] / a.A[i]
	}
}

//  Cast(
//    Vector3[int](0),
//    Vector3[float64](0),
//  ) -> Vector3[float64]{Vector3[int]~values}

func Cast[
	Value_b value.Values,
	Value_a value.Values,
	Length_a value.Arrays[Value_a],
	Length_b value.Arrays[Value_b],
](
	a Vector[Value_a, Length_a],
	b Vector[Value_b, Length_b],
) Vector[Value_b, Length_b] {
	var min_l int
	{
		var a_l Length_a
		var b_l Length_b
		if len(a_l) < len(b_l) {
			min_l = len(a_l)
		} else {
			min_l = len(b_l)
		}
	}
	for i := 0; i < min_l; i++ {
		b.A[i] = Value_b(a.A[i])
	}
	return b
}

func (vector *Vector[Value, Length]) Get(index int) Value {
	return vector.A[index]
}

func (vector *Vector[Value, Length]) Set(valuse Value, index int) {
	vector.A[index] = valuse
}

func (vector *Vector[Value, Length]) Len() int {
	return len(vector.A)
}

func (vector *Vector[Value, Length]) Mull_Scalar(a Vector[Value, Length]) (scalar Value) {
	for i := 0; i < len(vector.A); i++ {
		scalar += vector.A[i] * a.A[i]
	}
	return
}

/*
func (vector *Vector[Value, Length]) Mull_Matrix(a Vector[Value, Length])(scalar Value){
  for i:=0; i<len(vector.A);i++{
    scalar += vector.A[i] * a.A[i]
  }
  return
}
*/

/*

  Matrix3x3{
    X   X   X
    Y   Y   Y
    Z   Z   Z
  }

  Vectro3{
    A
    B
    C
  }

  Matrix3x3 X Vectro3 {

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

  }

*/
