package ds

import (
	"fmt"
	"slices"
	"strings"
)

type Number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
		~float64 | float32
}

type Matrix[T Number] [][]T

func NewMatrix[T Number](n int) Matrix[T] {
	mat := make([][]T, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]T, n)
	}
	return mat
}

func (mat Matrix[T]) String() string {
	var str strings.Builder

	str.WriteRune('\n')
	for _, row := range mat {
		fmt.Fprintf(&str, "%v\n", row)
	}

	return str.String()
}

func MatrixEqual[T Number](a Matrix[T], b Matrix[T]) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func MatrixClone[T Number](mat Matrix[T]) Matrix[T] {
	clone := make([][]T, len(mat))
	for i := 0; i < len(mat); i++ {
		clone[i] = slices.Clone(mat[i])
	}
	return clone
}
