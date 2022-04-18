package helpful_vector_matrix

import (
	value "github.com/illua1/go-helpful"
)

func Abc[Value value.Values](a Value)Value{
  if a >= 0 {
    return a
  } else {
    return -a
  }
}