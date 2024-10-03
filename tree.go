package main

import (
	"strconv"
)

type BinaryTreeNode[T any] struct {
	Data  T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func SerializeBinaryTree(root *BinaryTreeNode[int]) []string {
	stream := make([]string, 0, 1)

	serializeBinaryTreeNode(root, &stream)

	return stream
}

func serializeBinaryTreeNode(node *BinaryTreeNode[int], stream *[]string) {
	if node == nil {
		*stream = append(*stream, "null")
		return
	}
	n := strconv.FormatInt(int64(node.Data), 10)
	*stream = append(*stream, n)
	serializeBinaryTreeNode(node.Left, stream)
	serializeBinaryTreeNode(node.Right, stream)
}

func DeserializeBinaryTree(stream []string) *BinaryTreeNode[int] {
	var root *BinaryTreeNode[int]
	i := 0
	deserializeBinaryTreeNode(&root, stream, &i)

	return root
}

func deserializeBinaryTreeNode(node **BinaryTreeNode[int], stream []string, i *int) {
	if *i >= len(stream) {
		return
	}
	token := stream[*i]
	*i++
	if token == "null" {
		return
	}

	n, err := strconv.ParseInt(token, 10, strconv.IntSize)
	if err != nil {
		// this should never happen
		panic("failed to parse int from BinaryTreeNode[int]")
	}
	*node = &BinaryTreeNode[int]{
		Data: int(n),
	}
	deserializeBinaryTreeNode(&(*node).Left, stream, i)
	deserializeBinaryTreeNode(&(*node).Right, stream, i)
}

func IsSameBinaryTree[T comparable](a *BinaryTreeNode[T], b *BinaryTreeNode[T]) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Data != b.Data {
		return false
	}

	isLeftSubTreeSame := IsSameBinaryTree(a.Left, b.Left)
	isRighSubTreeSame := IsSameBinaryTree(a.Right, b.Right)

	return isLeftSubTreeSame && isRighSubTreeSame
}
