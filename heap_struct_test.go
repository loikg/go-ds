package ds_test

import (
	"reflect"
	"testing"

	"github.com/loikg/go-ds"
)

type pair struct {
	first  int
	second int
}

func (p pair) LessThan(other pair) bool {
	return p.second < other.second
}

func TestStructMaxHeap_Push(t *testing.T) {
	heap := ds.NewStructMaxHeap[pair]()

	heap.Push(pair{1, 1})
	assertNewMaxStruct(t, heap, pair{1, 1})
	assertLen(t, heap, 1)

	heap.Push(pair{2, 2})
	assertNewMaxStruct(t, heap, pair{2, 2})
	assertLen(t, heap, 2)

	heap.Push(pair{3, 1})
	assertNewMaxStruct(t, heap, pair{2, 2})
	assertLen(t, heap, 3)

	heap.Push(pair{4, 4})
	assertNewMaxStruct(t, heap, pair{4, 4})
	assertLen(t, heap, 4)
}

func TestStructMaxHeap_Pop(t *testing.T) {
	heap := ds.NewStructMaxHeap[pair]()

	heap.Push(pair{1, 0})
	heap.Push(pair{2, 1})
	heap.Push(pair{3, 2})

	top := heap.Pop()
	assertEqual(t, pair{3, 2}, top)
	assertLen(t, heap, 2)

	top = heap.Pop()
	assertEqual(t, pair{2, 1}, top)
	assertLen(t, heap, 1)

	top = heap.Pop()
	assertEqual(t, pair{1, 0}, top)
	assertLen(t, heap, 0)
}

func TestStructMinHeap_Push(t *testing.T) {
	heap := ds.NewStructMinHeap[pair]()

	heap.Push(pair{1, 1})
	assertNewMaxStruct(t, heap, pair{1, 1})
	assertLen(t, heap, 1)

	heap.Push(pair{2, 2})
	assertNewMaxStruct(t, heap, pair{1, 1})
	assertLen(t, heap, 2)

	heap.Push(pair{3, 3})
	assertNewMaxStruct(t, heap, pair{1, 1})
	assertLen(t, heap, 3)

	heap.Push(pair{4, 0})
	assertNewMaxStruct(t, heap, pair{4, 0})
	assertLen(t, heap, 4)
}

func TestStructMinHeap_Pop(t *testing.T) {
	heap := ds.NewStructMinHeap[pair]()

	heap.Push(pair{1, 0})
	heap.Push(pair{2, 1})
	heap.Push(pair{3, 2})

	top := heap.Pop()
	assertEqual(t, pair{1, 0}, top)
	assertLen(t, heap, 2)

	top = heap.Pop()
	assertEqual(t, pair{2, 1}, top)
	assertLen(t, heap, 1)

	top = heap.Pop()
	assertEqual(t, pair{3, 2}, top)
	assertLen(t, heap, 0)
}

func assertEqual[T any](t *testing.T, expected, val T) {
	t.Helper()
	if !reflect.DeepEqual(expected, val) {
		t.Fatalf("expected the maximum value of the heap to be: %v but got %v", expected, val)
	}
}

func assertNewMaxStruct[T any](t *testing.T, h ds.Heap[T], expected T) {
	t.Helper()

	top := h.Peek()

	assertEqual(t, expected, top)
}
