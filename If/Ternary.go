package helpful_ternary

// Simple Ternary operator.
func Ternary[T any](f bool, a, b T)T{
  if f {
    return a
  }else{
    return b
  }
}



