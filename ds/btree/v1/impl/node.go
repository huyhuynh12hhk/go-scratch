package btree

import (
	"cmp"
)

// Node of tree
type Node[K cmp.Ordered, V any] struct {
	Key    K
	Value  V
	Left   *Node[K, V]
	Right  *Node[K, V]
	height int
}


// Return height of a node
//
// Height value is the maximum number of 
// edges between that node and a leaf node
// 
// Leaf node height count as 1
func (n *Node[K, V]) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// BF is difference in height between right and left of subtrees
//
// BF = height(right) - height(left)
//
// - 0 : balance
//
// - positive : right heavy
//
// - negative : left heavy
func (n *Node[K, V]) BalanceFactor() int {
	if n == nil {
		return 0
	}

	return n.Right.Height() - n.Left.Height()
}

// Insert a node and self-balance
func (n *Node[K, V]) Insert(key K, value V) *Node[K, V] {
	if n == nil {
		// fmt.Printf("Start insert %v\n", key)
		return &Node[K, V]{
			Key:    key,
			Value:  value,
			height: 1,
		}
	}

	// Duplicate key -> Alter value
	// TODO: Optimize in case string?
	if n.Key == key {
		n.Value = value
		return n
	}

	if key < n.Key {
		// fmt.Printf("Insert %v to left\n", key)
		n.Left = n.Left.Insert(key, value)
	} else {
		// fmt.Printf("Insert %v to right\n", key)
		n.Right = n.Right.Insert(key, value)
	}

	n.height = max(n.Left.Height(), n.Right.Height()) + 1

	// fmt.Printf("%v checkout height %d\n",n.Key, n.height)

	return n.rebalance()
}

// Self balance implementation
func (n *Node[K, V]) rebalance() *Node[K, V] {
	switch {
	// Case Left-Left
	case n.BalanceFactor() < -1 && n.Left.BalanceFactor() == -1:
		// fmt.Println("> Rotate Right")
		return n.rotateRight()
		// Case Right-Right
	case n.BalanceFactor() > 1 && n.Right.BalanceFactor() == 1:
		// fmt.Println("> Rotate Left")
		return n.rotateLeft()
		// Case Left-Right
	case n.BalanceFactor() < -1 && n.Left.BalanceFactor() == 1:
		// fmt.Println("> Rotate Left Right")
		return n.rotateLeftRight()
		// Case Right-Left
	case n.BalanceFactor() > 1 && n.Right.BalanceFactor() == -1:
		// fmt.Println("> Rotate Right Left")
		return n.rotateRightLeft()
	}
	return n
}

func (n *Node[K, V]) rotateLeft() *Node[K, V] {
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	r.height = max(r.Left.Height(), r.Right.Height()) + 1
	return r
}

func (n *Node[K, V]) rotateRight() *Node[K, V] {
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	l.height = max(l.Left.Height(), l.Right.Height()) + 1
	return l
}

func (n *Node[K, V]) rotateRightLeft() *Node[K, V] {
	n.Right = n.Right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

func (n *Node[K, V]) rotateLeftRight() *Node[K, V] {
	n.Left = n.Left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

// Search value by key
func (n *Node[K, V]) Search(s K) (*Node[K, V], bool) {
	var rs Node[K, V]

	if n == nil {
		return &rs, false
	}

	switch {
	case s == n.Key:
		return n, true
	case s < n.Key:
		return n.Left.Search(s)
	default:
		return n.Right.Search(s)
	}
}
