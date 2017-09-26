package tree

import "fmt"

// Tree is the common interface for all tree implementations.
type Tree interface {
	// Serach finds the key in the tree and retuns the corresponding value.
	// If the key cannot be found nil will be returned
	//
	// It runs in O(h) time with h := height(tree)
	Search(int) interface{}
	// Insert inserts the given value at the key
	//
	// It runs in O(h) time with h := height(tree)
	Insert(int, interface{})
	// Delete deletes the value at the key.
	// If the key is not in the tree nothing happens
	//
	// It runs in O(h) time with h := height(tree)
	Delete(int)
}

type binaryTree struct {
	root *binaryNode
	size int
}

type binaryNode struct {
	key                 int
	value               interface{}
	left, right, parent *binaryNode
}

// NewBinaryTree returns a new, empty binary tree.
// Beause it is unbalanced the height is O(n) with n := #Nodes in the tree.
//
// Therefore all Operations can run in O(n) in the worst case.
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
		successor := t.Min(node.right)
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
//     / \		     /  \
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

// Min finds the node with the smallest key in the tree.
func (t *binaryTree) Min(n *binaryNode) *binaryNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Max finds the node with the gratest key in the tree.
func (t *binaryTree) Max(n *binaryNode) *binaryNode {
	for n.right != nil {
		n = n.right
	}
	return n
}

// Successor finds the node with the next gratest key in the tree.
func (t *binaryTree) Successor(node *binaryNode) *binaryNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return t.Min(node.right)
	}
	parent := node.parent
	for parent != nil && node == parent.right {
		node = parent
		parent = parent.parent
	}
	return parent
}

// Predecessor finds the node with the next smallest key in the tree.
func (t *binaryTree) Predecessor(node *binaryNode) *binaryNode {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return t.Max(node.left)
	}
	parent := node.parent
	for parent != nil && node == parent.left {
		node = parent
		parent = parent.parent
	}
	return parent
}

func (n *binaryNode) String() string {
	return fmt.Sprintf("k: %d,v: %v", n.key, n.value)
}
