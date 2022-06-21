package helpful_node

import (
	"fmt"
	"sync"

	ter "github.com/illua1/go-helpful/If"
)

type LNode[T any] struct {
	L, R    *LNode[T]
	Contain T
}

func NewLNode[T any](contain T) **LNode[T] {
	var node = LNode[T]{nil, nil, contain}
	var nodePointer = &node
	return &nodePointer
}

func (n LNode[T]) String() string {
	var ret string
	var lNode_p = &n
	for {
		if lNode_p != nil {
			ret += fmt.Sprint(
				ter.Ternary(lNode_p.R != nil, "<", ""),
				fmt.Sprint(lNode_p.Contain),
				ter.Ternary(lNode_p.L != nil, ">", ""),
			)
			lNode_p = lNode_p.L
		} else {
			break
		}
	}
	return ret
}

func merge_l[T any](node_r, node_l *LNode[T]) (ret *LNode[T], ok bool) {
	if node_r.L != nil {
		ret = node_r.L
		ok = true
	}
	node_r.L = node_l
	node_l.R = node_r
	return
}

func merge_r[T any](node_r, node_l *LNode[T]) (ret *LNode[T], ok bool) {
	if node_l.R != nil {
		ret = node_l.R
		ok = true
	}
	node_l.R = node_r
	node_r.L = node_l
	return
}

func NoNil[T any](lNode **LNode[T]) (*LNode[T], bool) {
	if lNode == nil {
		return nil, false
	}
	if (*lNode) == nil {
		return nil, false
	}
	return *lNode, true
}

func ToEndL[T any](lNode **LNode[T]) *LNode[T] {
	if l, ok := NoNil(lNode); ok {
		for ; l.L != nil; l = l.L {
		}
		return l
	}
	return nil
}

func ToEndR[T any](lNode **LNode[T]) *LNode[T] {
	if r, ok := NoNil(lNode); ok {
		for ; r.R != nil; r = r.R {
		}
		return r
	}
	return nil
}

func LenL[T any](lNode **LNode[T]) int {
	if l, ok := NoNil(lNode); ok {
		i := 1
		for ; l.L != nil; l = l.L {
			i++
		}
		return i
	}
	return 0
}

func LenR[T any](lNode **LNode[T]) int {
	if r, ok := NoNil(lNode); ok {
		i := 1
		for ; r.R != nil; r = r.R {
			i++
		}
		return i
	}
	return 0
}

func AppendEndL[T any](lNode **LNode[T], contain T) **LNode[T] {
	var node = NewLNode(contain)
	if end := ToEndL(lNode); end != nil {
		merge_l(end, *node)
	}
	return node
}

func AppendEndR[T any](lNode **LNode[T], contain T) **LNode[T] {
	var node = NewLNode(contain)
	if end := ToEndR(lNode); end != nil {
		merge_r(end, *node)
	}
	return node
}

func Append[T any](lNode **LNode[T], contain T) **LNode[T] {
	if node_l, ok := NoNil(lNode); ok {
		var node = NewLNode(contain)
		if another, ok := merge_l(node_l, *node); ok {
			if another_, _ := merge_r(*node, another); another_ != node_l {
				panic("go-helpful: NodeList : Addend : <current><new><another> -> current> != another || current != <another")
			}
		}
		return node
	}
	return NewLNode(contain)
}

/* // NO USE THIS
func(lNode *LNode[T])Del(){
  if (lNode.L != nil) && (lNode.R != nil) {
    lNode.R.L = lNode.L
    lNode.L.R = lNode.R
    lNode.L = nil
    lNode.R = nil
  } else{
    if lNode.L != nil {
      lNode.L.R = nil
      lNode.L = nil
    }
    if lNode.R != nil {
      lNode.R.L = nil
      lNode.R = nil
    }
  }
}
*/

func Del[T any](lNode **LNode[T]) {
	if *lNode != nil {
		if ((*lNode).L != nil) && ((*lNode).R != nil) {
			merge_l((*lNode).R, (*lNode).L)
			(*lNode).L = nil
			(*lNode).R = nil
		} else {
			if (*lNode).L != nil {
				(*lNode).L.R = nil
				(*lNode).L = nil
			} else {
				if (*lNode).R != nil {
					(*lNode).R.L = nil
					(*lNode).R = nil
				} else {
					(*lNode) = nil
				}
			}
		}
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

func ForI[T any](lNode **LNode[T], do func(index int, contain T)) {
	if *lNode != nil {
		var lNode_p *LNode[T] = *lNode
		var index int = 0
		for {
			do(index, lNode_p.Contain)
			if lNode_p.L != nil {
				lNode_p = lNode_p.L
			} else {
				return
			}
			index++
		}
	}
}

func ForParallel[T any](lNode **LNode[T], do func(index int, contain T)) {
	if *lNode != nil {
		var lNode_p *LNode[T] = *lNode
		var index int = 0
		var group sync.WaitGroup
		for {
			go func() {
				group.Add(1)
				do(index, lNode_p.Contain)
				group.Done()
			}()
			if lNode_p.L != nil {
				lNode_p = lNode_p.L
			} else {
				return
			}
			index++
		}
		group.Wait()
	}
}

func Len[T any](lNode **LNode[T]) int {
	var count int = 0
	if *lNode != nil {
		var lNode_p *LNode[T] = *lNode
		for {
			count++
			if lNode_p.L != nil {
				lNode_p = lNode_p.L
			} else {
				return count
			}
		}
	}
	return 0
}
