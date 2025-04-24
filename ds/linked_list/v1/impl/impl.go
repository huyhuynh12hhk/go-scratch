package linked_list

type LinkedList[T any] struct {
	head, tail *Node[T]
}
type Node[T any] struct {
	next *Node[T]
	val  T
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func FromArray[T any, A ~[]T](arr A) *LinkedList[T] {
	lst := LinkedList[T]{}
	for _, item := range arr {
		lst.Push(item)
	}

	return &lst
}

func (lst *LinkedList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &Node[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &Node[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *LinkedList[T]) Values() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *LinkedList[T]) Size() int {
	return len(lst.Values())
}

func (lst *LinkedList[T]) RemoveNode(node Node[T]) bool {
	if lst.head == &node {
		lst.head = lst.head.next
		return true
	}

	currentNode := lst.head
	for currentNode.next != nil && currentNode.next != &node {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		return false
	}

	currentNode.next = currentNode.next.next
	return true

}
