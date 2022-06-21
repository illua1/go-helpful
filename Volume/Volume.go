package helpful_volume

import (
	"fmt"
	value "github.com/illua1/go-helpful"
)

type Point[Value value.Values] struct {
	X, Y, Z Value
}

func NewPoint[Value value.Values](value Value) Point[Value] {
	return Point[Value]{
		X: value,
		Y: value,
		Z: value,
	}
}

func (p Point[Value]) String() string {
	return "(" + fmt.Sprint(p.X) + "," + fmt.Sprint(p.Y) + "," + fmt.Sprint(p.Z) + ")"
}

func (p Point[Value]) Add(q Point[Value]) Point[Value] {
	return Point[Value]{
		p.X + q.X,
		p.Y + q.Y,
		p.Z + q.Z,
	}
}

func (p Point[Value]) Sub(q Point[Value]) Point[Value] {
	return Point[Value]{
		p.X - q.X,
		p.Y - q.Y,
		p.Z - q.Z,
	}
}

func (p Point[Value]) Div(k Value) Point[Value] {
	return Point[Value]{
		p.X / k,
		p.Y / k,
		p.Z / k,
	}
}

func (p Point[Value]) Mul(k Value) Point[Value] {
	return Point[Value]{
		p.X * k,
		p.Y * k,
		p.Z * k,
	}
}

func (p Point[Value]) Eq(q Point[Value]) bool {
	return p == q
}

func (p Point[Value]) And(s Point[Value]) Box[Value] {
	return Box[Value]{
		Min: Point[Value]{p.X, p.Y, p.Z},
		Max: Point[Value]{p.X + s.X, p.Y + s.Y, p.Z + s.Z},
	}
}

func (p Point[Value]) In(b Box[Value]) bool {
	return ((b.Min.X <= p.X) && (b.Min.Y <= p.Y) && (b.Min.Z <= p.Z)) && ((b.Max.X >= p.X) && (b.Max.Y >= p.Y) && (b.Max.Z >= p.Z))
}

type Box[Value value.Values] struct {
	Min, Max Point[Value]
}

func NewBox[Value value.Values](min_x, min_y, min_z, max_x, max_y, max_z Value) Box[Value] {
	return Box[Value]{
		Min: Point[Value]{min_x, min_y, min_z},
		Max: Point[Value]{max_x, max_y, max_z},
	}
}

func (b Box[Value]) String() string {
	return fmt.Sprint(b.Min) + "-" + fmt.Sprint(b.Max)
}

func (b Box[Value]) Dx() Value {
	return b.Max.X - b.Min.X
}

func (b Box[Value]) Dy() Value {
	return b.Max.Y - b.Min.Y
}

func (b Box[Value]) Dz() Value {
	return b.Max.Z - b.Min.Z
}

func (b Box[Value]) Size() Point[Value] {
	return b.Max.Sub(b.Min)
}

func (b Box[Value]) Add(p Point[Value]) Box[Value] {
	return Box[Value]{b.Min.Add(p), b.Max.Add(p)}
}

func (b Box[Value]) Sub(p Point[Value]) Box[Value] {
	return Box[Value]{b.Min.Sub(p), b.Max.Sub(p)}
}

func (b Box[Value]) Mul(p Value) Box[Value] {
	return Box[Value]{b.Min.Mul(p), b.Max.Mul(p)}
}

func (b Box[Value]) Intersect(v Box[Value]) Box[Value] {
	if b.Min.X < v.Min.X {
		b.Min.X = v.Min.X
	}
	if b.Min.Y < v.Min.Y {
		b.Min.Y = v.Min.Y
	}
	if b.Min.Z < v.Min.Z {
		b.Min.Z = v.Min.Z
	}

	if b.Max.X > v.Max.X {
		b.Max.X = v.Max.X
	}
	if b.Max.Y > v.Max.Y {
		b.Max.Y = v.Max.Y
	}
	if b.Max.Z > v.Max.Z {
		b.Max.Z = v.Max.Z
	}
	return b
}

func (b Box[Value]) Union(v Box[Value]) Box[Value] {
	if b.Min.X > v.Min.X {
		b.Min.X = v.Min.X
	}
	if b.Min.Y > v.Min.Y {
		b.Min.Y = v.Min.Y
	}
	if b.Min.Z > v.Min.Z {
		b.Min.Z = v.Min.Z
	}

	if b.Max.X < v.Max.X {
		b.Max.X = v.Max.X
	}
	if b.Max.Y < v.Max.Y {
		b.Max.Y = v.Max.Y
	}
	if b.Max.Z < v.Max.Z {
		b.Max.Z = v.Max.Z
	}
	return b
}

func (b Box[Value]) Canon() Box[Value] {
	if b.Min.X < b.Max.X {
		b.Min.X, b.Max.X = b.Max.X, b.Min.X
	}
	if b.Min.Y < b.Max.Y {
		b.Min.Y, b.Max.Y = b.Max.Y, b.Min.Y
	}
	if b.Min.Z < b.Max.Z {
		b.Min.Z, b.Max.Z = b.Max.Z, b.Min.Z
	}
	return b
}

func (b Box[Value]) Centre() Point[Value] {
	return b.Min.Add(b.Max).Div(2)
}

func (b Box[Value]) Points() [8]Point[Value] {
	return [8]Point[Value]{
		b.Min,
		Point[Value]{b.Min.X, b.Min.Y, b.Max.Z},
		Point[Value]{b.Min.X, b.Max.Y, b.Max.Z},
		Point[Value]{b.Min.X, b.Max.Y, b.Min.Z},
		Point[Value]{b.Max.X, b.Min.Y, b.Min.Z},
		Point[Value]{b.Max.X, b.Min.Y, b.Max.Z},
		b.Max,
		Point[Value]{b.Max.X, b.Max.Y, b.Min.Z},
	}
}

func (b Box[Value]) Colise(a Box[Value]) bool {
	if in(a.Min.X, a.Max.X, b.Min.X, b.Max.X) {
		if in(a.Min.Y, a.Max.Y, b.Min.Y, b.Max.Y) {
			if in(a.Min.Z, a.Max.Z, b.Min.Z, b.Max.Z) {
				return true
			}
		}
	}
	return false
}

func in[Value value.Values](min, max, in_min, in_max Value) bool {
	if (min <= in_min) && (in_min <= max) {
		return true
	} else {
		return (min <= in_max) && (in_max <= max)
	}
	return false
}

func (b Box[Value]) FaceCentres() [6]Point[Value] {
	centre := b.Centre()
	return [6]Point[Value]{
		Point[Value]{centre.X, centre.Y, b.Min.Z},
		Point[Value]{centre.X, centre.Y, b.Max.Z},
		Point[Value]{centre.X, b.Min.Y, centre.Z},
		Point[Value]{centre.X, b.Max.Y, centre.Z},
		Point[Value]{b.Min.X, centre.Y, centre.Z},
		Point[Value]{b.Max.X, centre.Y, centre.Z},
	}
}

func (b Box[Value]) Edges() [12][2]int {
	return [12][2]int{
		[2]int{0, 1},
		[2]int{1, 2},
		[2]int{2, 3},
		[2]int{3, 0},
		//
		[2]int{4, 5},
		[2]int{5, 6},
		[2]int{6, 7},
		[2]int{7, 4},
		//
		[2]int{0, 4},
		[2]int{1, 5},
		[2]int{2, 6},
		[2]int{3, 7},
	}
}

func (b Box[Value]) FaceEdges() [6][4]int {
	return [6][4]int{
		[4]int{0, 4, 8, 9},
		[4]int{1, 5, 9, 10},
		[4]int{2, 6, 10, 11},
		[4]int{3, 7, 11, 9},
		[4]int{0, 1, 2, 3},
		[4]int{4, 5, 6, 7},
	}
}

func (b Box[Value]) FacePoints() [6][4]int {
	return [6][4]int{
		[4]int{0, 1, 2, 3},
		[4]int{4, 5, 6, 7},
		[4]int{0, 1, 4, 5},
		[4]int{1, 2, 6, 7},
		[4]int{2, 3, 6, 7},
		[4]int{3, 0, 7, 4},
	}
}

func (b Box[Value]) FaceArea() [6][2]Value {
	dx, dy, dz := b.Dx(), b.Dy(), b.Dz()
	dx, dy, dz = dx/2, dy/2, dz/2
	return [6][2]Value{
		[2]Value{dx, dy},
		[2]Value{-dx, -dy},
		[2]Value{dx, dz},
		[2]Value{-dx, -dz},
		[2]Value{dz, dy},
		[2]Value{-dz, -dy},
	}
}

const (
	TopFace = iota
	BottomFace
	LeftFace
	RigthFace
	FrontFace
	TailFace
)

func FaceOpposite(index int) int {
	var Opposites = [6]int{
		BottomFace,
		TopFace,
		RigthFace,
		LeftFace,
		TailFace,
		FrontFace,
	}
	return Opposites[index]
}

const (
	A_top_Edges = iota
	B_top_Edges
	C_top_Edges
	D_top_Edges
	A_bottom_Edges
	B_bottom_Edges
	C_bottom_Edges
	D_bottom_Edges
	A_centre_Edges
	B_centre_Edges
	C_centre_Edges
	D_centre_Edges
)

const (
	A_Point = iota
	B_Point
	C_Point
	D_Point
	E_Point
	F_Point
	G_Point
	H_Point
)

/*

  TopFace = {
    A_top_Edges = {
      A_Point
      B_Point
    }
    B_top_Edges = {
      B_Point
      C_Point
    }
    C_top_Edges = {
      C_Point
      D_Point
    }
    D_top_Edges = {
      D_Point
      A_Point
    }
  }

  BottomFace = {
    A_bottom_Edges = {
      E_Point
      F_Point
    }
    B_bottom_Edges = {
      F_Point
      G_Point
    }
    C_bottom_Edges = {
      G_Point
      H_Point
    }
    D_bottom_Edges = {
      H_Point
      E_Point
    }
  }

  LeftFace = {
    A_top_Edges
    A_centre_Edges
    A_bottom_Edges
    B_centre_Edges
  }

  FrontFace = {
    B_top_Edges
    B_centre_Edges
    B_bottom_Edges
    C_centre_Edges
  }

  RigthFace = {
    C_top_Edges
    C_centre_Edges
    C_bottom_Edges
    D_centre_Edges
  }

  TailFace = {
    D_top_Edges
    D_centre_Edges
    D_bottom_Edges
    A_centre_Edges
  }

  A_centre_Edges = {
    A_Point
    E_Point
  }

  B_centre_Edges = {
    B_Point
    F_Point
  }

  C_centre_Edges = {
    C_Point
    G_Point
  }

  D_centre_Edges = {
    D_Point
    H_Point
  }


*/

type BoxContainerFaces[T any] [6]T

type BoxContainerEdges[T any] [12]T

type BoxContainerPoints[T any] [8]T
