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

func TestCopyTo(t *testing.T) {
  
  var(
    list = make([]int, 100)
    value = 415
  )
  CopyTo(list, value)
  
  if len(list) != 100 {
    t.Error("CopyTo corrupt lenght")
  }
  
  for i := range list {
    if list[i] != value {
      t.Error("CopyTo corrupted value at index :", i)
    }
  }
}

func TestGetLast(t *testing.T) {
  
  var(
    list = []int{1,2,3,4,5,6,7,8,9}
    last = 9
  )
  
  if v, _ := GetLast(list, 14); v != last {
    t.Error("GetLast corrupt last at index: ", 14)
  }
  
  if v, _ := GetLast(list, 5); v  != list[5] {
    t.Error("GetLast corrupt last at index: ", 5)
  }
}

func TestGetFirst(t *testing.T) {
  var(
    list = []int{1,2,3,4,5,6,7,8,9}
    first = 1
  )
  
  if v, _ := GetFirst(list, -14); v != first {
    t.Error("GetFirst corrupt last at index: ", -14)
  }
  
  if v, _ := GetLast(list, 7); v  != list[7] {
    t.Error("GetFirst corrupt last at index: ", 7)
  }
}