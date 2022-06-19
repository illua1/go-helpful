package helpful_node

import (
	"testing"
)

func TestLNode(t *testing.T) {
	
  var node *LNode[int]
  
  Append(&node, 2)
  Append(&node, 7)
  Append(&node, 5)
  var delNode = Append(&node, 267)
  Append(&node, 27)
  Append(&node, 2)
  
  var first = node.String()
  
  delNode.Del()
  
  var dual = node.String()
  
  if first == dual {
    t.Error(first, " == ", dual, " After Del()")
  }
}
