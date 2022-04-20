package helpful_map

import(
  "testing"
)

const count = 100

func BenchmarkMyMap(b *testing.B) {
  
  var my_map = NewMap[int, int]()
  
  for i := 0; i < count; i++ {
    my_map.Set(i,i)
  }
  for i := 0; i < count; i++ {
    if v, _ := my_map.Find(i); v != i {}
  }
}

func BenchmarkMap(b *testing.B) {
  
  var _map = map[int]int{}
  
  for i := 0; i < count; i++ {
    _map[i] = i
  }
  for i := 0; i < count; i++ {
    if v, _ := _map[i]; v != i {}
  }
}