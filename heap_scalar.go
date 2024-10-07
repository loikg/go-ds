package ds

import (
	"cmp"
)

type scalarHeapContainer[T cmp.Ordered] struct {
	arr    []T
	lesser func(T, T) bool
}

func (h scalarHeapContainer[T]) Len() int           { return len(h.arr) }
func (h scalarHeapContainer[T]) Less(i, j int) bool { return h.lesser(h.arr[i], h.arr[j]) }
func (h scalarHeapContainer[T]) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *scalarHeapContainer[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.arr = append(h.arr, x.(T))
}

func (h *scalarHeapContainer[T]) Pop() any {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}

func (h scalarHeapContainer[T]) Peek() any {
	return h.arr[0]
}
