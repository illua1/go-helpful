package helpful_vector_matrix

import (
	//"math"
	//"image"
	"testing"
	//value "github.com/illua1/go-helpful"
	compare "github.com/illua1/go-helpful/Compare"
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

  Collums{
    Row
    Row
    Row
  }

  Matrix[0] -> Row[0] -> X

  Matrix[X][Y]

*/

func TestNewMatrix(t *testing.T) {

	matrixMax := Matrix6x6[int]()
	{
		m := Matrix5x5[int]()
		{
			ms := matrixMax.Slise(0, 0, 5, 5)
			if !MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix5x5 corrupted\n", ms, "\n", m)
			}
		}
		{
			ms := matrixMax.Slise(0, 1, 5, 6)
			if MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix5x5 corrupted\n", ms, "\n", m)
			}
		}
	}
	{
		m := Matrix4x4[int]()
		{
			ms := matrixMax.Slise(0, 0, 4, 4)
			if !MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix4x4 corrupted\n", ms, "\n", m)
			}
		}
		{
			ms := matrixMax.Slise(0, 1, 4, 5)
			if MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix4x4 corrupted\n", ms, "\n", m)
			}
		}
	}
	{
		m := Matrix3x3[int]()
		{
			ms := matrixMax.Slise(0, 0, 3, 3)
			if !MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix3x3 corrupted\n", ms, "\n", m)
			}
		}
		{
			ms := matrixMax.Slise(0, 1, 3, 4)
			if MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix3x3 corrupted\n", ms, "\n", m)
			}
		}
	}
	{
		m := Matrix2x2[int]()
		{
			ms := matrixMax.Slise(0, 0, 2, 2)
			if !MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix2x2 corrupted\n", ms, "\n", m)
			}
		}
		{
			ms := matrixMax.Slise(0, 1, 2, 3)
			if MatrixIsEqual[int, int](ms, &m) {
				t.Error("Matrix2x2 corrupted\n", ms, "\n", m)
			}
		}
	}
}

func TestMatrixMultiplication(t *testing.T) {
	m1 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 4, 7},
			[3]int{2, 5, 8},
			[3]int{3, 6, 9},
		},
	}
	m2 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{9, 6, 3},
			[3]int{8, 5, 2},
			[3]int{7, 4, 1},
		},
	}
	m3 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{30, 24, 18},
			[3]int{84, 69, 54},
			[3]int{138, 114, 90},
		},
	}
	if m1.Mull(m2) != m3 {
		t.Error("Matrix (3x3) multiplication corrupt\n", m1, "\n", m2, "\n", m3, "\n", m1.Mull(m2))
	}
}

func TestMatrixTransposed(t *testing.T) {
	m1 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 4, 7},
			[3]int{2, 5, 8},
			[3]int{3, 6, 9},
		},
	}
	m2 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 2, 3},
			[3]int{4, 5, 6},
			[3]int{7, 8, 9},
		},
	}
	if m1.Transposed() != m2 {
		t.Error("Matrix (3x3) transposed corrupt\n", m1, "\n", m2, "\n", m1.Transposed())
	}
}

func TestMatrixMinor(t *testing.T) {
	m := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 4, 7},
			[3]int{2, 5, 8},
			[3]int{3, 6, 9},
		},
	}
	m1 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{5, 8, 0},
			[3]int{6, 9, 0},
			[3]int{0, 0, 0},
		},
	}
	m2 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 7, 0},
			[3]int{3, 9, 0},
			[3]int{0, 0, 0},
		},
	}
	m3 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 4, 0},
			[3]int{2, 5, 0},
			[3]int{0, 0, 0},
		},
	}
	m4 := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{4, 7, 0},
			[3]int{5, 8, 0},
			[3]int{0, 0, 0},
		},
	}
	if *(m.Minor(0, 0).(*Matrix[int, [3]int, [3][3]int])) != m1 {
		t.Error(m, m1, m.Minor(0, 0), 0, 0)
	}
	if *(m.Minor(1, 1).(*Matrix[int, [3]int, [3][3]int])) != m2 {
		t.Error(m, m2, m.Minor(1, 1), 1, 1)
	}
	if *(m.Minor(2, 2).(*Matrix[int, [3]int, [3][3]int])) != m3 {
		t.Error(m, m3, m.Minor(2, 2), 2, 0)
	}
	if *(m.Minor(2, 0).(*Matrix[int, [3]int, [3][3]int])) != m4 {
		t.Error(m, m4, m.Minor(2, 0), 2, 0)
	}
}

func TestMatrixMinorMutable(t *testing.T) {
	m_ := Matrix[int, [3]int, [3][3]int]{
		[3][3]int{
			[3]int{1, 4, 7},
			[3]int{2, 5, 8},
			[3]int{3, 6, 9},
		},
	}
	m := m_.Mutable()
	m1_ := Matrix[int, [2]int, [2][2]int]{
		[2][2]int{
			[2]int{5, 8},
			[2]int{6, 9},
		},
	}
	m1 := m1_.Mutable()
	m2_ := Matrix[int, [2]int, [2][2]int]{
		[2][2]int{
			[2]int{1, 7},
			[2]int{3, 9},
		},
	}
	m2 := m2_.Mutable()
	m3_ := Matrix[int, [2]int, [2][2]int]{
		[2][2]int{
			[2]int{1, 4},
			[2]int{2, 5},
		},
	}
	m3 := m3_.Mutable()
	m4_ := Matrix[int, [2]int, [2][2]int]{
		[2][2]int{
			[2]int{4, 7},
			[2]int{5, 8},
		},
	}
	m4 := m4_.Mutable()
	if !compare.Compare(m.Minor(0, 0).Body(), m1.Body()) {
		t.Error(m.Body(), m1.Body(), m.Minor(0, 0).Body(), "> ", 0, 0)
	}
	if !compare.Compare(m.Minor(1, 1).Body(), m2.Body()) {
		t.Error(m.Body(), m2.Body(), m.Minor(1, 1).Body(), "> ", 1, 1)
	}
	if !compare.Compare(m.Minor(2, 2).Body(), m3.Body()) {
		t.Error(m.Body(), m3.Body(), m.Minor(2, 2).Body(), "> ", 2, 0)
	}
	if !compare.Compare(m.Minor(2, 0).Body(), m4.Body()) {
		t.Error(m.Body(), m4.Body(), m.Minor(2, 0).Body(), "> ", 2, 0)
	}
}
