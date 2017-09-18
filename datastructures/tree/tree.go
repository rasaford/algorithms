package tree

import (
	"fmt"
	"strings"

	"github.com/rasaford/algorithms/datastructures/stack"
)

type Tree interface {
	String() string
}

// binaryTree is a container for a binary tree
type binaryTree struct {
	root *binaryNode
}

// binaryNode is a container for a node in a binary tree
type binaryNode struct {
	value       int
	left, right *binaryNode
}

// NewBinaryTree returns a new, empty binary tree
func NewBinaryTree() Tree {
	return &binaryTree{}
}

// String retuns all the values in the binary tree concatenated together..
func (t *binaryTree) String() string {
	out := ""
	do := func(a int) {
		out = fmt.Sprintf("%s %d", out, a)
	}
	t.root.walkIterative(do)
	return strings.TrimSpace(out)
}

// walkRecursive performs a recursive depth first search on the subtree at the given
// node and exectues the function do() for each value.
//
// It runs in O(n) time with  n := len(tree)
// It uses O(1) extra space in memory
func (n *binaryNode) walkRecursive(do func(int)) {
	if n == nil {
		return
	}
	do(n.value)
	n.left.walkRecursive(do)
	n.right.walkRecursive(do)
}

// walkIterative performs a iterative depth first search on the subtree at the given
// node and exectues the function do() for each value.
//
// It runs in O(n) time with  n := len(tree)
// It uses O(n) extra space in memory
func (n *binaryNode) walkIterative(do func(int)) {
	if n == nil {
		return
	}
	stack := stack.New()
	for {
		do(n.value)
		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
		if node, err := stack.Pop(); err != nil {
			break
		} else {
			n = node.(*binaryNode)
		}
	}
}
