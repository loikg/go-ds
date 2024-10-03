package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Push(elem T) {
	s.arr = append(s.arr, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.arr) < 1 {
		return *new(T), false
	}
	length := len(s.arr)
	elem := s.arr[length-1]
	s.arr = slices.Delete(s.arr, length-1, length)
	return elem, true
}

func (s Stack[T]) Peek() (T, bool) {
	if len(s.arr) < 1 {
		return *new(T), false
	}

	return s.arr[len(s.arr)-1], true
}

func (s Stack[T]) Len() int {
	return len(s.arr)
}

func (s Stack[T]) String() string {
	var str strings.Builder

	str.WriteRune('[')
	for i, elem := range s.arr {
		if i != 0 {
			str.WriteRune(',')
		}
		if char, ok := any(elem).(rune); ok {
			fmt.Fprintf(&str, "%v", strconv.QuoteRune(char))
		} else {
			fmt.Fprintf(&str, "%v", elem)
		}
	}
	str.WriteRune(']')

	return str.String()
}
