package helpful_slise

import(
  "github.com/illua1/go-helpful"
)

// Cast will do type cast each elements of input slise to second type.
func Cast[A, B helpful.Values](a []A) (b []B) {
	b = make([]B, len(a))
	for i := range a {
		b[i] = B(a[i])
	}
	return
}

// Fill will do rewrite input slise by function output.
func Fill[T any](in []T, method func(index int) T) {
	for i := range in {
		in[i] = method(i)
	}
}

// Join return slise that will contained all element of input sliset, order will saved.
func Join[T any](slises ...[]T) (out []T) {
	var length int = 0
	for i := range slises {
		length += len(slises[i])
	}
	out = make([]T, length)
	length = 0
	for i := range slises {
		copy(
			out[length:length+len(slises[i])],
			slises[i],
		)
		length += len(slises[i])
	}
	return
}

// CopyTo will fast filling slise by input value.
func CopyTo[T any](slise []T, value T) {
	l := len(slise)
	if l < 3 {
		for i := 0; i < l; i++ {
			slise[i] = value
		}
		return
	}
	slise[0], slise[1] = value, value
	size := 2
	for {
		copy(slise[size:], slise[:size])
		size *= 2
		if size >= l {
			break
		}
	}
}

func GetLast[T any](in []T, index int)(T, bool){
  if len(in) > 0 {
    if (len(in) > index) && (index >= 0) {
      return in[index], true
    } else {
      return in[len(in)-1], true
    }
  } else {
    var v T
    return v, false 
  }
}

func GetFirst[T any](in []T, index int)(T, bool){
  if len(in) > 0 {
    if (len(in) > index) && (index >= 0) {
      return in[index], true
    } else {
      return in[0], true
    }
  } else {
    var v T
    return v, false 
  }
}