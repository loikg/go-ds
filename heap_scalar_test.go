package ds_test

import (
	"cmp"
	"testing"

	"github.com/loikg/go-ds"
)

func TestScalarMaxHeap_Push(t *testing.T) {
	heap := ds.NewScalarMaxHeap[int]()

	heap.Push(1)
	assertNewMaxScalar(t, heap, 1)
	assertLen(t, heap, 1)

	heap.Push(2)
	assertNewMaxScalar(t, heap, 2)
	assertLen(t, heap, 2)

	heap.Push(3)
	assertNewMaxScalar(t, heap, 3)
	assertLen(t, heap, 3)

	heap.Push(4)
	assertNewMaxScalar(t, heap, 4)
	assertLen(t, heap, 4)

	heap.Push(-1)
	assertNewMaxScalar(t, heap, 4)
	assertLen(t, heap, 5)
}

func TestScalarMaxHeap_Pop(t *testing.T) {
	heap := ds.NewScalarMaxHeap[int]()

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)

	top := heap.Pop()
	assertEqual(t, 3, top)
	assertLen(t, heap, 2)

	top = heap.Pop()
	assertEqual(t, 2, top)
	assertLen(t, heap, 1)

	top = heap.Pop()
	assertEqual(t, 1, top)
	assertLen(t, heap, 0)
}

func TestScalarMinHeap_Push(t *testing.T) {
	heap := ds.NewScalarMinHeap[int]()

	heap.Push(1)
	assertNewMaxScalar(t, heap, 1)
	assertLen(t, heap, 1)

	heap.Push(2)
	assertNewMaxScalar(t, heap, 1)
	assertLen(t, heap, 2)

	heap.Push(-1)
	assertNewMaxScalar(t, heap, -1)
	assertLen(t, heap, 3)
}

func TestScalarMinHeap_Pop(t *testing.T) {
	heap := ds.NewScalarMinHeap[int]()

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	top := heap.Pop()
	assertEqual(t, 1, top)
	assertLen(t, heap, 2)

	top = heap.Pop()
	assertEqual(t, 2, top)
	assertLen(t, heap, 1)

	top = heap.Pop()
	assertEqual(t, 3, top)
	assertLen(t, heap, 0)
}

type comparableAndOrdered interface {
	comparable
	cmp.Ordered
}

func assertLen[T any](t *testing.T, h ds.Heap[T], expected int) {
	t.Helper()
	if l := h.Len(); l != expected {
		t.Errorf("expected len %d but got %d", expected, l)
	}
}

func assertNewMaxScalar[T comparableAndOrdered](t *testing.T, h ds.Heap[T], expected T) {
	t.Helper()

	maxInt := h.Peek()

	if maxInt != expected {
		t.Fatalf("expected the maximum value of the heap to be: %v but got %v", expected, maxInt)
	}
}
