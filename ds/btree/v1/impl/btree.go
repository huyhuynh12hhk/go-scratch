package btree

import (
	"cmp"
	"strings"
)


// Binary Tree type
type BinaryTree[K cmp.Ordered, V any] struct {
	Root *Node[K, V]
}

// Create new blank binary tree
func NewBinaryTree[K cmp.Ordered, V any]() *BinaryTree[K, V] {
	return &BinaryTree[K, V]{}
}

// Create binary tree from array of values
func FromValueArray[A ~[]V, V cmp.Ordered](items A) *BinaryTree[V, V]{
	bt:= BinaryTree[V, V]{}
	for _, v := range items{
		bt.Insert(v, v)
	}

	return &bt
}

// Create binary tree from a string
func FromAString(str string) *BinaryTree[string, string]{
	bt:= BinaryTree[string, string]{}
	for _, v := range strings.Split(str, " "){
		bt.Insert(v, v)
	}

	return &bt
}

// Insert key value pair to tree
func (t *BinaryTree[K, V]) Insert(key K, value V) {
	t.Root = t.Root.Insert(key, value)
	if t.Root.BalanceFactor() < -1 || t.Root.BalanceFactor() > 1 {
		t.rebalance()
	}
}

// Rebalance by tree
func (t *BinaryTree[K, V]) rebalance() {
	if t == nil || t.Root == nil {
		return
	}
	t.Root = t.Root.rebalance()
}

// Search value by key from tree
func (t *BinaryTree[K, V]) Search(key K) (*Node[K,V], bool) {
	var rs Node[K, V]
	if t == nil || t.Root == nil {
		return &rs, false
	}
	return t.Root.Search(key)
}

// Return an array of nodes in in-order traversal form
func (t *BinaryTree[K, V]) ToInOrderTraversalNodeArray() []Node[K,V]{

	arr:= []Node[K,V]{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		walk(n.Left, depth+1)
		arr= append(arr, *n)
		walk(n.Right, depth+1)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of nodes in pre-order traversal form
func (t *BinaryTree[K, V]) ToPreOrderTraversalNodeArray() []Node[K,V]{

	arr:= []Node[K,V]{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		arr= append(arr, *n)
		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of nodes in post-order traversal form
func (t *BinaryTree[K, V]) ToPostOrderTraversalNodeArray() []Node[K,V]{

	arr:= []Node[K,V]{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
		arr= append(arr, *n)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of values in in-order traversal form
func (t *BinaryTree[K, V]) ToInOrderTraversalValues() []V{

	arr:= []V{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		walk(n.Left, depth+1)
		arr= append(arr, n.Value)
		walk(n.Right, depth+1)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of keys in pre-order traversal form
func (t *BinaryTree[K, V]) ToPreOrderTraversalKeys() []K{

	arr:= []K{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		arr= append(arr, n.Key)
		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of keys in in-order traversal form
func (t *BinaryTree[K, V]) ToInOrderTraversalKeys() []K{

	arr:= []K{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		walk(n.Left, depth+1)
		arr= append(arr, n.Key)
		walk(n.Right, depth+1)
	}

	walk(t.Root, 0)

	return arr
}

// Return an array of keys in post-order traversal form
func (t *BinaryTree[K, V]) ToPostOrderTraversalKeys() []K{

	arr:= []K{}
	// A local func for recursive
	var walk func(*Node[K, V], int)
	walk = func(n *Node[K, V], depth int) {
		if n == nil {
			return
		}
		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
		arr= append(arr, n.Key)
	}

	walk(t.Root, 0)

	return arr
}