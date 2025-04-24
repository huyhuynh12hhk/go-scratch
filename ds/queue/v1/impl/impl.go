package queue

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func FromArray[T any, A ~[]T](arr A) *Queue[T] {
	qu := Queue[T]{}
	for _, item := range arr {
		qu.queue = append(qu.queue, item)
	}

	return &qu
}

func (qu *Queue[T]) Enqueue(item T) {
	qu.queue = append(qu.queue, item)
}

func (qu *Queue[T]) Dequeue() (T, bool) {
	var item T
	if !qu.IsEmpty() {
		item, qu.queue = qu.queue[0], qu.queue[1:]
		return item, true
	}
	return item, false

}

func (qu *Queue[T]) Peek() (T, bool) {
	var item T
	if !qu.IsEmpty() {
		item = qu.queue[0]
		return item, true
	}
	return item, false
}

func (qu *Queue[T]) Size() int {
	return len(qu.queue)
}

func (qu *Queue[T]) IsEmpty() bool {
	return len(qu.queue) == 0
}
