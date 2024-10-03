package ds

import "slices"

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Enqueue(elem T) {
	q.arr = append(q.arr, elem)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.arr) < 1 {
		return *new(T), false
	}
	elem := q.arr[0]
	q.arr = slices.Delete(q.arr, 0, 1)

	return elem, true
}

func (q Queue[T]) Len() int {
	return len(q.arr)
}

func (q Queue[T]) Peek() (T, bool) {
	if len(q.arr) < 1 {
		return *new(T), false
	}
	return q.arr[0], true
}
