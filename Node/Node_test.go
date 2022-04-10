package helpful_node

import (
  "testing"
  slise "github.com/illua1/go-helpful/Slise"
  "math/rand"
)

func newBNodeWood(size int)*BNode[int]{
  list := make([]BNode[int], size)
  slise.Fill(
    list, 
    func(index int)BNode[int]{
      return NewBNode(nil, nil, rand.Int())
    },
  )
  for i := 1; i < size; i++{
    BNodeDescrent(
      &list[0], 
      &list[i], 
      func(a,b *BNode[int])bool{
        return a.Contain > b.Contain
      },
    )
  }
  return &list[0]
}

func TestBNodeSize(t *testing.T) {
  var node = newBNodeWood(1000)
  if BNodeSize(node) != 1000 {
    t.Error("BNode size corrupt: ", 1000, ",", BNodeSize(node))
  }
}