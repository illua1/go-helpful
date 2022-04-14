package helpful_node

import (
	"fmt"
	ter "github.com/illua1/go-helpful/If"
)

type BNode[T any] struct {
	L, R    *BNode[T]
	Contain T
}

func (n BNode[T]) String() string {
	return "{" + fmt.Sprint(
		ter.Ternary(n.L != nil, fmt.Sprint(n.L), ""),
		", ",
		fmt.Sprint(n.Contain),
		", ",
		ter.Ternary(n.R != nil, fmt.Sprint(n.R), ""),
	) + "}"
}

func NewBNode[T any](l, r *BNode[T], t T) BNode[T] {
	return BNode[T]{l, r, t}
}

func BNodeDescrent[T any](root, descrent *BNode[T], condition func(a, b *BNode[T]) bool) {
	root_ := root
	for {
		if condition(root_, descrent) {
			if root_.R != nil {
				root_ = root_.R
				continue
			} else {
				root_.R = descrent
				return
			}
		} else {
			if root_.L != nil {
				root_ = root_.L
				continue
			} else {
				root_.L = descrent
				return
			}
		}
	}
}

func BNodeLengthL[T any](branch *BNode[T]) int {
	l := branch
	counter := 0
	for {
		if l != nil {
			l = l.L
			counter++
		} else {
			return counter
		}
	}
	return 0
}

func BNodeLengthR[T any](branch *BNode[T]) int {
	r := branch
	counter := 0
	for {
		if r != nil {
			r = r.R
			counter++
		} else {
			return counter
		}
	}
	return 0
}

func BNodeSize[T any](branch *BNode[T]) (r int) {
	if branch.R != nil {
		r += BNodeSize[T](branch.R)
	}
	if branch.L != nil {
		r += BNodeSize[T](branch.L)
	}
	r += 1
	return
}

func BNodeGetL[T any](branch *BNode[T]) (ret *BNode[T]) {
	ret = branch.L
	if ret == nil {
		ret = branch
		return
	}
	for {
		if ret.L != nil {
			ret = ret.L
		} else {
			return
		}
	}
}

func BNodeGetR[T any](branch *BNode[T]) (ret *BNode[T]) {
	ret = branch.R
	if ret == nil {
		ret = branch
		return
	}
	for {
		if ret.R != nil {
			ret = ret.R
		} else {
			return
		}
	}
}

func BNodeFillTo[T any](node *BNode[T], slise []T) {
	if len(slise) != BNodeSize(node) {
		return
	}
	pointer := 0
	bNodeFull[T](node, &pointer, slise)
}

func bNodeFull[T any](node *BNode[T], pointer *int, slise []T) {
	if node.L != nil {
		bNodeFull[T](node.L, pointer, slise)
	}
	slise[*pointer] = node.Contain
	*pointer += 1
	if node.R != nil {
		bNodeFull[T](node.R, pointer, slise)
	}
}

func BNodeForTo[T any](node *BNode[T], method func(index int, contain T)) {
	pointer := 0
	bNodeFor[T](node, &pointer, method)
}

func bNodeFor[T any](node *BNode[T], pointer *int, method func(index int, contain T)) {
	if node.L != nil {
		bNodeFor[T](node.L, pointer, method)
	}
	method(*pointer, node.Contain)
	*pointer += 1
	if node.R != nil {
		bNodeFor[T](node.R, pointer, method)
	}
}
