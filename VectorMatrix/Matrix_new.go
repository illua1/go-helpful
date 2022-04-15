package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
)

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
