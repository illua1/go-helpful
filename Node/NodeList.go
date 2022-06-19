package helpful_node

import (
	"fmt"
)

type LNode[T any] struct {
	L, R    *LNode[T]
	Contain T
}

func (n LNode[T]) String() string {
  var ret string
  ret += fmt.Sprint(n.Contain)
  var lNode_p *LNode[T] = n.L
  for {
    if lNode_p != nil {
      ret += ", "+fmt.Sprint(lNode_p.Contain)
      lNode_p = lNode_p.L
    } else{
      break
    }
  }
  return ret
}

func Append[T any](lNode **LNode[T], contain T)*LNode[T]{
  if *lNode != nil {
    var lNode_p *LNode[T] = *lNode
    for {
      if lNode_p.L == nil {
        lNode_p.L = &LNode[T]{nil, lNode_p, contain}
        return lNode_p.L
      } else {
        lNode_p = lNode_p.L
      }
    }
  } else {
    *lNode = &LNode[T]{nil, nil, contain}
    return *lNode
  }
  return nil
}

func(lNode *LNode[T])Del(){
  if (lNode.L != nil) && (lNode.R != nil) {
    lNode.R.L = lNode.L
    lNode.L.R = lNode.R
    lNode.L = nil
    lNode.R = nil
  } else{
    
  }
}

func For[T any](lNode **LNode[T], do func(contain T)) {
  if *lNode != nil {
    var lNode_p *LNode[T] = *lNode
    for {
      do(lNode_p.Contain)
      if lNode_p.L != nil {
        lNode_p = lNode_p.L
      } else {
        return
      }
    }
  }
}