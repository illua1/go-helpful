package helpful_vector_matrix

import(
  "math"
  value "github.com/illua1/go-helpful"
)

type Matrix[
  Value value.Values,
  Row value.Arrays[Value],
  Collums value.Arrays[Row],
] struct {
  A Collums
}

func Matrix3x3[Value value.Values]()Matrix[Value, [3]Value, [3][3]Value]{
  var m Matrix[Value, [3]Value, [3][3]Value]
  m.A[0][0] = 1
  m.A[1][1] = 1
  m.A[2][2] = 1
  return m
}

func(matrix Matrix[Value, Row, Collums])Mull(B Matrix[Value, Row, Collums])(C Matrix[Value, Row, Collums]){
  var len_ int
  if len(matrix.A) > len(matrix.A[0]) {
    len_ = len(matrix.A[0])
  }else{
    len_ = len(matrix.A)
  }
  for row := 0; row < len_; row++ {
    for coll := 0; coll < len_; coll++ {
      for i := 0; i < len_; i++ {
        C.A[row][coll] += matrix.A[row][i] * B.A[coll][i]
      }
    }
  }
  return
}
/*
func(matrix Matrix[Value, Row, Collums])Mull_Vector(V Vector[Value, Row])(C Matrix[Value, Row, Collums]){

  return
}
*/
func Rotate3x3[Value value.Values, Angle value.Values](angle [3]Angle)Matrix[Value, [3]Value, [3][3]Value]{
  var m = Matrix3x3[Value]()
  
  m = m.Mull(Rotate3x3_x[Value](angle[0]))
  m = m.Mull(Rotate3x3_y[Value](angle[1]))
  m = m.Mull(Rotate3x3_z[Value](angle[2]))
  
  return m
}

func Rotate3x3_x[Value value.Values, Angle value.Values](angle Angle)Matrix[Value, [3]Value, [3][3]Value]{
  var m Matrix[Value, [3]Value, [3][3]Value]
  m.A[0][0] = 1
  m.A[1][0] = Value(math.Cos(float64(angle)))
  m.A[1][1] = Value(-math.Sin(float64(angle)))
  m.A[2][0] = Value(math.Sin(float64(angle)))
  m.A[2][1] = Value(math.Cos(float64(angle)))
  return m
}

func Rotate3x3_y[Value value.Values, Angle value.Values](angle Angle)Matrix[Value, [3]Value, [3][3]Value]{
  var m Matrix[Value, [3]Value, [3][3]Value]
  m.A[1][1] = 1
  m.A[0][0] = Value(math.Cos(float64(angle)))
  m.A[0][2] = Value(math.Sin(float64(angle)))
  m.A[2][0] = Value(-math.Sin(float64(angle)))
  m.A[2][2] = Value(math.Cos(float64(angle)))
  return m
}

func Rotate3x3_z[Value value.Values, Angle value.Values](angle Angle)Matrix[Value, [3]Value, [3][3]Value]{
  var m Matrix[Value, [3]Value, [3][3]Value]
  m.A[2][2] = 1
  m.A[0][0] = Value(math.Cos(float64(angle)))
  m.A[0][1] = Value(-math.Sin(float64(angle)))
  m.A[1][0] = Value(math.Sin(float64(angle)))
  m.A[1][1] = Value(math.Cos(float64(angle)))
  return m
}