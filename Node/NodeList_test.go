package helpful_node

import (
	"testing"
)

func TestLNode(t *testing.T) {

	var node = NewLNode(1)
	Append(node, 2)
	Append(node, 7)
	Append(node, 5)
	var delNode = Append(node, 267)
	Append(delNode, 11)
	Append(node, 27)
	Append(node, 2)
	var first = (*node).String()
	Del(delNode)
	var dual = (*node).String()

	if first == dual {
		t.Error(first, " == ", dual, " After Del(...)")
	}
}
