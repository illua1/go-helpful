package helpful_slise

import (
	"runtime"
	"testing"
)

const (
	count   = 100
	lengths = 10000
)

func BenchmarkJoin(b *testing.B) {
	b.StopTimer()
	list := make([][]int, count)
	Fill(list, func(index int) []int { return make([]int, lengths) })
	b.StartTimer()

	Join(list...)

	list = nil
	runtime.GC()
}

func BenchmarkJoinSimpleAppend(b *testing.B) {
	b.StopTimer()
	list := make([][]int, count)
	Fill(list, func(index int) []int { return make([]int, lengths) })
	b.StartTimer()

	var unioned []int
	for i := range list {
		unioned = append(unioned, list[i]...)
	}

	unioned = nil
	runtime.GC()
}

func BenchmarkCopyToFast(b *testing.B) {
	b.StopTimer()
	list := make([]int, 1000000)
	b.StartTimer()
	CopyTo(list, 54)
}

func BenchmarkCopyToSimple(b *testing.B) {
	b.StopTimer()
	list := make([]int, 1000000)
	b.StartTimer()
	for i := 0; i < len(list); i++ {
		list[i] = 54
	}
}
