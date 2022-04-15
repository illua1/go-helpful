package helpful_vector_matrix

import (
	"math/rand"
	//"image"
	"testing"
	value "github.com/illua1/go-helpful"
)

func Check[Value value.Values](a, b, c Matrix_Main_Functions[Value], t *testing.T, str string){
  if !MatrixIsEqual[Value, Value](a, b) {
    t.Error("Corrupted [", str, "] :\n", a, "\n not a:\n", b)
  }
  if !MatrixIsEqual[Value, Value](a, c) {
    t.Error("Corrupted [", str, "] :\n", a, "\n not a:\n", c)
  }
  if !MatrixIsEqual[Value, Value](c, b) {
    t.Error("Corrupted [", str, "] :\n", c, "\n not a:\n", b)
  }
}

type Tester func(in Matrix_Main_Functions[int])Matrix_Main_Functions[int]

func TestMatrixInterfaces(t *testing.T) {
  
  var do = []Tester{
    func(in Matrix_Main_Functions[int])Matrix_Main_Functions[int]{
      return in.Minor(0,0)
    },
  }
  var texts = []string{
    "Minor at 3,3 for 6x6 matrix",
  }
  var constant_size = [][2]int{
    [2]int{5, 5},
  }
  
  
  contain := Matrix6x6[int]()
  contain.FillAs(func(x, y int)int{
    return rand.Intn(100)
  })
  
  for i := range do {
    var (
      a = Matrix6x6[int]()
      b_p = Matrix6x6[int]()
      b = (&b_p).Slise(0,0,6,6)
      c = a.Mutable()
    )
    MatrixWrite[int, int](&a, &contain)
    MatrixWrite[int, int](b, &contain)
    MatrixWrite[int, int](&c, &contain)
    Check[int](
      do[i](&a).Slise(0,0, constant_size[i][0], constant_size[i][1]),
      do[i](b),
      do[i](&c),
      t,
      texts[i],
    )
  }
}