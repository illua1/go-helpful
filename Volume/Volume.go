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

func (p Point[Value]) And(s Point[Value]) Boxe[Value] {
	return Boxe[Value]{
		Min: Point[Value]{p.X, p.Y, p.Z},
		Max: Point[Value]{p.X + s.X, p.Y + s.Y, p.Z + s.Z},
	}
}

func (p Point[Value]) In(b Boxe[Value]) bool {
	return ((b.Min.X <= p.X) && (b.Min.Y <= p.Y) && (b.Min.Z <= p.Z)) && ((b.Max.X >= p.X) && (b.Max.Y >= p.Y) && (b.Max.Z >= p.Z))
}

type Boxe[Value value.Values] struct {
	Min, Max Point[Value]
}

func NewBoxe[Value value.Values](min_x, min_y, min_z, max_x, max_y, max_z Value) Boxe[Value] {
	return Boxe[Value]{
		Min: Point[Value]{min_x, min_y, min_z},
		Max: Point[Value]{max_x, max_y, max_z},
	}
}

func (b Boxe[Value]) String() string {
	return fmt.Sprint(b.Min) + "-" + fmt.Sprint(b.Max)
}

func (b Boxe[Value]) Dx() Value {
	return b.Max.X - b.Min.X
}

func (b Boxe[Value]) Dy() Value {
	return b.Max.Y - b.Min.Y
}

func (b Boxe[Value]) Dz() Value {
	return b.Max.Z - b.Min.Z
}

func (b Boxe[Value]) Size() Point[Value] {
	return b.Max.Sub(b.Min)
}

func (b Boxe[Value]) Add(p Point[Value]) Boxe[Value] {
	return Boxe[Value]{b.Min.Add(p), b.Max.Add(p)}
}

func (b Boxe[Value]) Sub(p Point[Value]) Boxe[Value] {
	return Boxe[Value]{b.Min.Sub(p), b.Max.Sub(p)}
}

func (b Boxe[Value]) Mul(p Value) Boxe[Value] {
	return Boxe[Value]{b.Min.Mul(p), b.Max.Mul(p)}
}

func (b Boxe[Value]) Intersect(v Boxe[Value]) Boxe[Value] {
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

func (b Boxe[Value]) Union(v Boxe[Value]) Boxe[Value] {
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

func (b Boxe[Value]) Canon() Boxe[Value] {
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

func (b Boxe[Value]) Centre() Point[Value] {
  return b.Min.Add(b.Max).Div(2)
}

func (b Boxe[Value]) Points() [8]Point[Value] {
	return [8]Point[Value]{
		b.Min,
		Point[Value]{b.Min.X, b.Min.Y, b.Max.Z},
		Point[Value]{b.Min.X, b.Max.Y, b.Min.Z},
		Point[Value]{b.Min.X, b.Max.Y, b.Max.Z},
		Point[Value]{b.Max.X, b.Min.Y, b.Min.Z},
		Point[Value]{b.Max.X, b.Min.Y, b.Max.Z},
		Point[Value]{b.Max.X, b.Max.Y, b.Min.Z},
		b.Max,
	}
}

func (b Boxe[Value]) FaceCentres() [6]Point[Value]{
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

func (b Boxe[Value]) Edges() [12][2]int {
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
		[2]int{0, 7},
		[2]int{1, 6},
		[2]int{2, 5},
		[2]int{3, 4},
	}
}

func (b Boxe[Value]) FaceEdges() [6][4]int {
	return [6][4]int{
		[4]int{0, 4, 8, 9},
		[4]int{1, 5, 9, 10},
		[4]int{2, 6, 10, 11},
		[4]int{3, 7, 11, 9},
		[4]int{0, 1, 2, 3},
		[4]int{4, 5, 6, 7},
	}
}

func (b Boxe[Value]) FacePoints() [6][4]int {
	return [6][4]int{
		[4]int{0, 1, 2, 3},
		[4]int{4, 5, 6, 7},
		[4]int{0, 1, 4, 5},
		[4]int{1, 2, 6, 7},
		[4]int{2, 3, 6, 7},
		[4]int{3, 0, 7, 4},
	}
}
/*
func (b Boxe[Value]) FaceCentres() (ret [6]Point[Value]){
	points := b.Points()
  f_indeces := b.FacePoints()
  for i := range f_indeces {
    for p := range f_indeces[i] {
      ret[i] = ret[i].Add(points[f_indeces[i][p]])
    }
    ret[i] = ret[i].Div(4)
  }
  return
}
*/