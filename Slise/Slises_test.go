package helpful_slise

import (
  "testing"
)

func TestJoin(t *testing.T) {
  var(
    a = []int{1,2,3,4,5,6,7,8,9,0}
    b = []int{0,1,2,5,89,27,356,2}
    c = []int{1,2,3,4,5,6,7,8,9,0,0,1,2,5,89,27,356,2}
  )
  if len(Join(a,b)) != len(a)+len(b){
    t.Error("Different length after Join(", a, "and", b, ") to ", c)
  }
}

func TestFill(t *testing.T) {
  var(
    list []int = make([]int, 100)
  )
  
  Fill(list, func(index int)int{
    return index
  })
  
  if len(list) != 100{
    t.Error("Fill corrupted size length")
  }
  
  for i := 0; i < 100; i++{
    if list[i] != i {
      t.Error("Fill element at index ", i, "corrupted: ", list[i])
    }
  }
}

func TestCast(t *testing.T) {
  
  var(
    a = []int{1,2,3,4,100,1000,50000}
    b = []float64{1.,2.,3.,4.,100.,1000.,50000.}
    c = Cast[int, float64](a)
  )
  
  if (len(a) != len(b)) || (len(a) != len(c)) {
    t.Error("Cast int to float64 is corrupted al lenght :", a, b, c)
  }else{
    for i := range c {
      if c[i] != b[i] {
        t.Error("Cast int to float64 is corrupted at index", i, ":", a[i], "->", c[i])
      }
    }
  }
}