package helpful_slise

import(
  "testing"
)

const(
  count = 1000
  lengths = 100000
)

func BenchmarkJoin(b *testing.B) {
  b.StopTimer()
  list := make([][]int, count)
  Fill(list, func(index int)[]int{return make([]int, lengths)})
  b.StartTimer()
  
  Join(list...)
}

func BenchmarkJoinSimpleAppend(b *testing.B) {
  b.StopTimer()
  list := make([][]int, count)
  Fill(list, func(index int)[]int{return make([]int, lengths)})
  b.StartTimer()
  
  var unioned []int
  for i := range list {
    unioned = append(unioned, list[i]...)
  }
}