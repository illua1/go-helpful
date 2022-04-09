package helpful_sort

import(
  "testing"
  "fmt"
)

func TestSortMax(t *testing.T){
  
  list := [][]int{
    []int{4,7,8,3,9,5,3,6,3,7,2,2,73,88,2,7,4,6,87,2,37,13,17,36},
    []int{4,7},
    []int{4},
  }
  for i := range list {
    SortMax[int](list[i])
    fmt.Println(list[i])
  }
}