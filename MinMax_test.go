package helpful

import(
  "testing"
)

type F_Test[T Values] struct{
  a, b, c T
}

func TestMaxF(t *testing.T) {
  var(
    list = []F_Test[int]{
      F_Test[int]{1,2,2},
      F_Test[int]{5,2,5},
      F_Test[int]{1,21,21},
      F_Test[int]{144,211,211},
    }
  )

  for i := range list{
    if MaxF(list[i].a, list[i].b) != list[i].c {
      t.Error("MaxF at",[]int{list[i].a, list[i].b}," is not a", list[i].c)
    }
  }
}

func TestMax(t *testing.T){
  var(
    list = [][]int{
      []int{1,2,3,4,5,67,8,89,0,1544},
      []int{555,222,888,111,000,-11,55},
      []int{654,557,157875,156,11},
    }
    max = []int{
      1544,
      888,
      157875,
    }
  )
  for i := range list {
    if Max(list[i]...) != max[i] {
      t.Error("Max at", list[i], " is not a",max[i])
    }
  }
}

func TestMinF(t *testing.T) {
  var(
    list = []F_Test[int]{
      F_Test[int]{1,2,1},
      F_Test[int]{5,2,2},
      F_Test[int]{1,21,1},
      F_Test[int]{144,211,144},
    }
  )

  for i := range list{
    if MinF(list[i].a, list[i].b) != list[i].c {
      t.Error("MinF at",[]int{list[i].a, list[i].b}," is not a", list[i].c)
    }
  }
}

func TestMin(t *testing.T){
  var(
    list = [][]int{
      []int{1,2,3,4,5,67,8,89,0,1544},
      []int{555,222,888,111,000,-11,55},
      []int{654,557,157875,156,11},
    }
    min = []int{
      0,
      -11,
      11,
    }
  )
  for i := range list {
    if Min(list[i]...) != min[i] {
      t.Error("Min at", list[i], " is not a",min[i])
    }
  }
}


func TestMaxIdF(t *testing.T) {
  var(
    list = []F_Test[int]{
      F_Test[int]{1,2,1},
      F_Test[int]{5,2,0},
      F_Test[int]{1,21,1},
      F_Test[int]{144,211,1},
    }
  )

  for i := range list{
    if MaxIdF(list[i].a, list[i].b) != list[i].c {
      t.Error("MaxIdF at",[]int{list[i].a, list[i].b}," is not at index", list[i].c)
    }
  }
}

func TestMaxId(t *testing.T){
  var(
    list = [][]int{
      []int{1,2,3,4,5,67,8,89,0,1544},
      []int{555,222,888,111,000,-11,55},
      []int{654,557,157875,156,11},
    }
    max = []int{
      9,
      2,
      2,
    }
  )
  for i := range list {
    if MaxId(list[i]...) != max[i] {
      t.Error("MaxId at", list[i], " is not at index",max[i])
    }
  }
}

func TestMinIdF(t *testing.T) {
  var(
    list = []F_Test[int]{
      F_Test[int]{1,2,0},
      F_Test[int]{5,2,1},
      F_Test[int]{1,21,0},
      F_Test[int]{144,211,0},
    }
  )

  for i := range list{
    if MinIdF(list[i].a, list[i].b) != list[i].c {
      t.Error("MinIdF at",[]int{list[i].a, list[i].b}," is not at index", list[i].c)
    }
  }
}

func TestMinId(t *testing.T){
  var(
    list = [][]int{
      []int{1,2,3,4,5,67,8,89,0,1544},
      []int{555,222,888,111,000,-11,55},
      []int{654,557,157875,156,11},
    }
    min = []int{
      8,
      5,
      4,
    }
  )
  for i := range list {
    if MinId(list[i]...) != min[i] {
      t.Error("MinId at", list[i], " is not at index",min[i])
    }
  }
}