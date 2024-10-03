package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSerializeTree(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		root     *BinaryTreeNode[int]
		expected []string
	}{
		{
			name:     "given a nil root return null",
			root:     nil,
			expected: []string{"null"},
		},
		{
			name: "given a tree with left and right nodes",
			root: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 2,
				},
				Right: &BinaryTreeNode[int]{
					Data: 5,
					Left: &BinaryTreeNode[int]{
						Data: 6,
					},
					Right: &BinaryTreeNode[int]{
						Data: 7,
					},
				},
			},
			expected: []string{"3", "2", "null", "null", "5", "6", "null", "null", "7", "null", "null"},
		},
		{
			name: "given a tree with right nodes only",
			root: &BinaryTreeNode[int]{
				Data: 3,
				Right: &BinaryTreeNode[int]{
					Data: 5,
					Right: &BinaryTreeNode[int]{
						Data: 7,
						Right: &BinaryTreeNode[int]{
							Data: 9,
						},
					},
				},
			},
			expected: []string{"3", "null", "5", "null", "7", "null", "9", "null", "null"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := SerializeBinaryTree(tc.root)
			if !slices.Equal(tc.expected, actual) {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestDeserializeBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		stream   []string
		expected *BinaryTreeNode[int]
	}{
		{
			name:     "given a nil root return null",
			stream:   []string{"null"},
			expected: nil,
		},
		{
			name:   "given a tree with left and right nodes",
			stream: []string{"3", "2", "null", "null", "5", "6", "null", "null", "7", "null", "null"},
			expected: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 2,
				},
				Right: &BinaryTreeNode[int]{
					Data: 5,
					Left: &BinaryTreeNode[int]{
						Data: 6,
					},
					Right: &BinaryTreeNode[int]{
						Data: 7,
					},
				},
			},
		},
		{
			name:   "given a tree with right nodes only",
			stream: []string{"3", "null", "5", "null", "7", "null", "9", "null", "null"},
			expected: &BinaryTreeNode[int]{
				Data: 3,
				Right: &BinaryTreeNode[int]{
					Data: 5,
					Right: &BinaryTreeNode[int]{
						Data: 7,
						Right: &BinaryTreeNode[int]{
							Data: 9,
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := DeserializeBinaryTree(tc.stream)
			if !IsSameBinaryTree(tc.expected, actual) {
				t.Errorf(
					"expected %v but got %v",
					SerializeBinaryTree(tc.expected),
					SerializeBinaryTree(actual),
				)
			}
		})
	}
}

func TestIsSameBinaryTree(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a        *BinaryTreeNode[int]
		b        *BinaryTreeNode[int]
		expected bool
	}{
		{
			a: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 9,
					Left: &BinaryTreeNode[int]{
						Data: 10,
					},
					Right: &BinaryTreeNode[int]{
						Data: 5,
					},
				},
				Right: &BinaryTreeNode[int]{
					Data: 7,
				},
			},
			b: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 9,
					Left: &BinaryTreeNode[int]{
						Data: 10,
					},
					Right: &BinaryTreeNode[int]{
						Data: 5,
					},
				},
				Right: &BinaryTreeNode[int]{
					Data: 7,
				},
			},
			expected: true,
		},
		{
			a: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 9,
					Left: &BinaryTreeNode[int]{
						Data: 10,
					},
					Right: &BinaryTreeNode[int]{
						Data: 5,
					},
				},
				Right: &BinaryTreeNode[int]{
					Data: 7,
				},
			},
			b: &BinaryTreeNode[int]{
				Data: 3,
				Left: &BinaryTreeNode[int]{
					Data: 9,
					Left: &BinaryTreeNode[int]{
						Data: 10,
					},
					Right: &BinaryTreeNode[int]{
						Data: 5,
					},
				},
				Right: &BinaryTreeNode[int]{
					Data: 8,
				},
			},
			expected: false,
		},
		{
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			a: nil,
			b: &BinaryTreeNode[int]{
				Data: 0,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		name := fmt.Sprintf(
			"given %v and %v expect %v",
			SerializeBinaryTree(tc.a),
			SerializeBinaryTree(tc.b),
			tc.expected,
		)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := IsSameBinaryTree(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
