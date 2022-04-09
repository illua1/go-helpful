package helpful_sort

import(
  "testing"
  "math/rand"
  slise "github.com/illua1/go-helpful/Slise"
)

const size = 100000

func BenchmarkSortMax(b *testing.B) {
  b.StopTimer()
  list := make([]int, size)
  slise.Fill(list, func(index int)int{return rand.Int()})
  b.StartTimer()
  
  SortMax(list)
}

func BenchmarkSortAny(b *testing.B) {
  b.StopTimer()
  list := make([]int, size)
  slise.Fill(list, func(index int)int{return rand.Int()})
  b.StartTimer()
  
  SortAny(list, func(a, b int)bool{return a > b})
}

func BenchmarkSortAnyFloat(b *testing.B) {
  b.StopTimer()
  list := make([]float64, size)
  slise.Fill(list, func(index int)float64{return rand.Float64()})
  b.StartTimer()
  
  SortAny(list, func(a, b float64)bool{return a > b})
}