package helpful_sort

import (
	slise "github.com/illua1/go-helpful/Slise"
	"math/rand"
	"testing"
)

func SortCheck[T any](slise_ []T, condition func(a, b T) bool, t *testing.T, name string) {
	for i := 0; i < len(slise_)-1; i++ {
		if !condition(slise_[i], slise_[i+1]) {
			t.Error("Sort", name, "corrupted at index :", i)
			break
		}
	}
}

func TestSortMax(t *testing.T) {
	list := make([]int, 10000)
	for i := 0; i < 100; i++ {
		slise.Fill(list, func(index int) int { return rand.Int() })
		SortMax[int](list)
		SortCheck(list, func(a, b int) bool { return a >= b }, t, "Max")
	}
}

func TestSortMin(t *testing.T) {
	list := make([]int, 10000)
	for i := 0; i < 100; i++ {
		slise.Fill(list, func(index int) int { return rand.Int() })
		SortMin[int](list)
		SortCheck(list, func(a, b int) bool { return a <= b }, t, "Min")
	}
}

func TestSortAny(t *testing.T) {
	list := make([]int, 10000)
	for i := 0; i < 100; i++ {
		slise.Fill(list, func(index int) int { return rand.Int() })
		SortAny[int](list, func(a, b int) bool { return a > b })
		SortCheck(list, func(a, b int) bool { return a >= b }, t, "Any(>)")
	}
}
