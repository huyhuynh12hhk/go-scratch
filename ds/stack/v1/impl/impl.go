package stack

type Stack[T any] struct{
	items [] T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func FromArray[T any, A ~[]T](arr A) *Stack[T] {
	stk := Stack[T]{}
	for _, item := range arr {
		stk.items = append(stk.items, item)
	}

	return &stk
}


func (stk *Stack[T]) Push(item T){
	stk.items = append(stk.items, item)
}

func (stk *Stack[T]) Pop()(T, bool){
	var item T
	if !stk.IsEmpty(){
		item, stk.items = stk.items[len(stk.items)-1], stk.items[:len(stk.items)-1] 
		return item, true
	}

	return item, false
}

func (stk *Stack[T]) Peek()(T, bool){
	var item T
	if !stk.IsEmpty() {
		item = stk.items[len(stk.items)-1]
		return item, true
	}
	return item, false
}

func (stk *Stack[T]) IsEmpty() bool{
	return len(stk.items) == 0
}

func (stk *Stack[T]) Size() int{
	return len(stk.items)
}
