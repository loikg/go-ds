package main

import (
	"fmt"
	"strings"
)

type LinkListNode[T any] struct {
	Data T
	Next *LinkListNode[T]
}

func (node *LinkListNode[T]) String() string {
	var str strings.Builder
	visited := make(map[*LinkListNode[T]]struct{})

	str.WriteRune('[')

	current := node
	for current != nil {
		if _, ok := visited[current]; ok {
			str.WriteString("...")
			break
		}
		visited[current] = struct{}{}
		fmt.Fprintf(&str, "%v", current.Data)
		if current.Next != nil {
			str.WriteRune(',')
		}
		current = current.Next
	}
	str.WriteRune(']')

	return str.String()
}

func NewLinkListFromSlice[T any](data []T) *LinkListNode[T] {
	if len(data) < 1 {
		return nil
	}

	head := &LinkListNode[T]{
		Data: data[0],
		Next: nil,
	}

	current := head
	for _, d := range data[1:] {
		current.Next = &LinkListNode[T]{
			Data: d,
			Next: nil,
		}
		current = current.Next
	}

	return head
}

func LinklistEqual[T comparable](list1 *LinkListNode[T], list2 *LinkListNode[T]) bool {
	current1 := list1
	current2 := list2

	for current1 != nil && current2 != nil {
		if current1.Data != current2.Data {
			return false
		}
		current1 = current1.Next
		current2 = current2.Next
	}

	return current1 == nil && current2 == nil
}
