package helpful_sort

import (
	"github.com/illua1/go-helpful"
	node "github.com/illua1/go-helpful/Node"
	slise "github.com/illua1/go-helpful/Slise"
)

func SortMax[T helpful.Values](in []T) {
	l := len(in)

	if l < 2 {
		return
	}

	Wood := make([]node.BNode[T], l)

	slise.Fill(
		Wood,
		func(index int) node.BNode[T] {
			return node.NewBNode(nil, nil, in[index])
		},
	)

	for i := 1; i < l; i++ {
		node.BNodeDescrent(
			&Wood[0],
			&Wood[i],
			func(a, b *node.BNode[T]) bool {
				return a.Contain > b.Contain
			},
		)
	}

	node.BNodeFillTo[T](&Wood[0], in)
}

func SortMin[T helpful.Values](in []T) {
	l := len(in)

	if l < 2 {
		return
	}

	Wood := make([]node.BNode[T], l)

	slise.Fill(
		Wood,
		func(index int) node.BNode[T] {
			return node.NewBNode(nil, nil, in[index])
		},
	)

	for i := 1; i < l; i++ {
		node.BNodeDescrent(
			&Wood[0],
			&Wood[i],
			func(a, b *node.BNode[T]) bool {
				return a.Contain < b.Contain
			},
		)
	}

	node.BNodeFillTo[T](&Wood[0], in)
}

func SortAny[T any](in []T, condition func(a, b T) bool) {
	l := len(in)

	if l < 2 {
		return
	}

	Wood := make([]node.BNode[T], l)

	slise.Fill(
		Wood,
		func(index int) node.BNode[T] {
			return node.NewBNode(nil, nil, in[index])
		},
	)

	for i := 1; i < l; i++ {
		node.BNodeDescrent(
			&Wood[0],
			&Wood[i],
			func(a, b *node.BNode[T]) bool {
				return condition(a.Contain, b.Contain)
			},
		)
	}

	node.BNodeFillTo[T](&Wood[0], in)
}
