package ds

import (
	"cmp"
	"container/heap"
)

type Heap[T any] interface {
	Push(T)
	Pop() T
	Peek() T
	Len() int
	Empty() bool
}

type peekerHeap interface {
	heap.Interface
	Peek() any
}

type scalarHeapWrapper[T cmp.Ordered] struct {
	heap peekerHeap
}

func (h *scalarHeapWrapper[T]) Push(val T) {
	heap.Push(h.heap, val)
}

func (h *scalarHeapWrapper[T]) Pop() T {
	return heap.Pop(h.heap).(T)
}

func (h *scalarHeapWrapper[T]) Peek() T {
	return h.heap.Peek().(T)
}

func (h scalarHeapWrapper[T]) Empty() bool {
	return h.heap.Len() == 0
}

func (h scalarHeapWrapper[T]) Len() int {
	return h.heap.Len()
}

func NewScalarMaxHeap[T cmp.Ordered]() Heap[T] {
	h := &scalarHeapContainer[T]{
		lesser: func(t1, t2 T) bool {
			return t1 > t2
		},
	}
	heap.Init(h)
	return &scalarHeapWrapper[T]{h}
}

func NewScalarMinHeap[T cmp.Ordered]() Heap[T] {
	h := &scalarHeapContainer[T]{
		lesser: func(t1, t2 T) bool {
			return t1 < t2
		},
	}
	heap.Init(h)
	return &scalarHeapWrapper[T]{h}
}

type structHeapWrapper[T Ordered[T]] struct {
	heap peekerHeap
}

func (h *structHeapWrapper[T]) Push(val T) {
	heap.Push(h.heap, val)
}

func (h *structHeapWrapper[T]) Pop() T {
	return heap.Pop(h.heap).(T)
}

func (h *structHeapWrapper[T]) Peek() T {
	return h.heap.Peek().(T)
}

func (h structHeapWrapper[T]) Empty() bool {
	return h.heap.Len() == 0
}

func (h structHeapWrapper[T]) Len() int {
	return h.heap.Len()
}

func NewStructMinHeap[T Ordered[T]]() Heap[T] {
	h := &StructHeap[T]{
		lesser: func(t1, t2 T) bool {
			return t1.LessThan(t2)
		},
	}
	heap.Init(h)
	return &structHeapWrapper[T]{h}
}

func NewStructMaxHeap[T Ordered[T]]() Heap[T] {
	h := &StructHeap[T]{
		lesser: func(t1, t2 T) bool {
			return t2.LessThan(t1)
		},
	}
	heap.Init(h)
	return &structHeapWrapper[T]{h}
}
