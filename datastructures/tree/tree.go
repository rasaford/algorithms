package tree

import "fmt"

type Tree interface {
	Search(int) interface{}
	Insert(int, interface{})
	Delete(int)
	predecessor(*binaryNode) *binaryNode
	successor(*binaryNode) *binaryNode
	min(*binaryNode) *binaryNode
	max(*binaryNode) *binaryNode
}

// binaryTree is a container for a binary search tree.
type binaryTree struct {
	root *binaryNode
	size int
}

// binaryNode is a container for a node in a BST.
type binaryNode struct {
	key                 int
	value               interface{}
	left, right, parent *binaryNode
}

// NewBinaryTree returns a new, empty binary tree
func NewBinaryTree() Tree {
	return &binaryTree{}
}

func (t *binaryTree) Search(key int) interface{} {
	node := t.search(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *binaryTree) search(key int) *binaryNode {
	n := t.root
	for n != nil && key != n.key {
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func (t *binaryTree) Insert(key int, value interface{}) {
	insert := &binaryNode{
		key:   key,
		value: value,
	}
	// maintains a training pointer to get back to the parent.
	node, trailing := t.root, t.root
	for node != nil {
		trailing = node
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	insert.parent = trailing
	if trailing == nil {
		t.root = insert
	} else if key < trailing.key {
		trailing.left = insert
	} else {
		trailing.right = insert
	}
	t.size++
}

// Delete deletes a key from the tree.
//
// It runs in O(h) time with h := height(tree)
func (t *binaryTree) Delete(key int) {
	node := t.search(key)
	if node == nil {
		return
	}
	// when there is no left or right child is nil the current node can be deleted
	// by replacing it with the other child node.
	if node.left == nil {
		t.transplant(node, node.right)
	} else if node.right == nil {
		t.transplant(node, node.left)
	} else {
		successor := t.min(node.right)
		// if successor is node's right child replace it with it's right child.
		// (the left child has to be nil)
		if successor.parent != node {
			t.transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = node
		}
		// replace node with successor
		t.transplant(node, successor)
		successor.left = node.left
		successor.left.parent = node
	}
	t.size--
}

// transplant replaces u with v in the tree. The node u is discarded.
//    p			      p
//  /  \			/  \
//      u	->         v
//     / \		      / \
//    v
//  /  \
func (t *binaryTree) transplant(u, v *binaryNode) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// min finds the node with the smallest key in the tree.
func (t *binaryTree) min(n *binaryNode) *binaryNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

// max finds the node with the gratest key in the tree.
func (t *binaryTree) max(n *binaryNode) *binaryNode {
	for n.right != nil {
		n = n.right
	}
	return n
}

// successor finds the node with the next gratest key in the tree.
func (t *binaryTree) successor(node *binaryNode) *binaryNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return t.min(node.right)
	}
	parent := node.parent
	for parent != nil && node == parent.right {
		node = parent
		parent = parent.parent
	}
	return parent
}

// predecessor finds the node with the next smallest key in the tree.
func (t *binaryTree) predecessor(node *binaryNode) *binaryNode {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return t.max(node.left)
	}
	parent := node.parent
	for parent != nil && node == parent.left {
		node = parent
		parent = parent.parent
	}
	return parent
}

// String retuns all the values in the binary tree concatenated together..
// func (t *binaryTree) String() string {
// 	out := ""
// 	do := func(a int) {
// 		out = fmt.Sprintf("%s %d", out, a)
// 	}
// 	t.root.walkIterative(do)
// 	return strings.TrimSpace(out)
// }

// walkRecursive performs a recursive depth first search on the subtree at the given
// node and exectues the function do() for each value.
//
// It runs in O(n) time with  n := len(tree)
// It uses O(1) extra space in memory
// func (n *binaryNode) walkRecursive(do func(int)) {
// 	if n == nil {
// 		return
// 	}
// 	do(n.value)
// 	n.left.walkRecursive(do)
// 	n.right.walkRecursive(do)
// }

// walkIterative performs a iterative depth first search on the subtree at the given
// node and exectues the function do() for each value.
//
// It runs in O(n) time with  n := len(tree)
// It uses O(n) extra space in memory
// func (n *binaryNode) walkIterative(do func(int)) {
// 	if n == nil {
// 		return
// 	}
// 	stack := stack.New()
// 	for {
// 		do(n.value)
// 		if n.right != nil {
// 			stack.Push(n.right)
// 		}
// 		if n.left != nil {
// 			stack.Push(n.left)
// 		}
// 		if node, err := stack.Pop(); err != nil {
// 			break
// 		} else {
// 			n = node.(*binaryNode)
// 		}
// 	}
// }

func (n *binaryNode) String() string {
	return fmt.Sprintf("k: %d,v: %v", n.key, n.value)
}
