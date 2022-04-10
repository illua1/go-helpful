package helpful

import (
	"golang.org/x/exp/constraints"
)

// Main helpful math type interface.
type Values interface {
	constraints.Float | constraints.Integer | byte
}


type array_read[T any] interface {
  Get(index int)T
  First()T
  Last()T
}

type array_write[T any] interface {
  Set(index int, value T)
  Fill(value T)
}

type array_resize[T any] interface {
  Append(start_end bool, value T)int
  Join(start_end bool, values ...[]T)int
  Connect(start_end bool, values ...Range[T])int
  
  Delet(index int)(ok bool, size int)
  Mask(mask []bool)int
}

type Iterator[T any] interface {
  Loop()(value T, index int, ok bool)
}

type Range[T any] interface {
  array_read[T]
  array_write[T]
  
  Len()int
  End()int
  
  GStart()int
  GIndex(index int)int
  
  IndexValid(index int)bool
  
  Slise(offset int, length int)Range[T]
  SliseD(point1 float64, point2 float64)Range[T]
  
  Equal(b Range[T])bool
  
  CopyTy(b Range[T])bool
  
  Iterat()Iterator[T]
}

type Slise[T any] interface {
  Range[T]
  array_resize[T]
}