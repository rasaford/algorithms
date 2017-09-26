package tree

import (
	"fmt"
	"math"

	"github.com/rasaford/algorithms/internal/helper"
)

// Properties of a Red-Black-Tree:
// 1. Every node is either red or black
// 2. The root is black
// 3. Every leaf is black
// 4. If a node is red then both it's children are black
// 5. For each node, all simple paths from the node to descendant leaves contain the
// same number of black nodes.

type rbTree struct {
	root     *rbNode
	sentinel *rbNode
	size     int
}

type rbNode struct {
	key                 int
	value               interface{}
	left, right, parent *rbNode
	// color of the node in the RB-Tree. true is red, false is black
	color bool
}

// NewRedBlackTree creates an empty Red - Black tree. A RB-Tree is a balanced tree
// that guarantees that it's height is always O(lg n), with n := #Nodes in the tree.
//
// Therefore all Operations on it run in O(lg n) time.
func NewRedBlackTree() Tree {
	s := &rbNode{
		key:   math.MinInt32,
		color: false, // all leafs have to be black
	}
	return &rbTree{
		root:     s,
		sentinel: s,
	}
}

func (t *rbTree) Search(key int) interface{} {
	node := t.search(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (t *rbTree) search(key int) *rbNode {
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

func (t *rbTree) Insert(key int, value interface{}) {
	insert := &rbNode{
		key:   key,
		value: value,
		color: true, // all inserted nodes are initially colored red
		left:  t.sentinel,
		right: t.sentinel,
	}
	// maintains a training pointer to get back to the parent.
	trailing, node := t.sentinel, t.root
	for node != t.sentinel {
		trailing = node
		if insert.key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	insert.parent = trailing
	if trailing == t.sentinel {
		t.root = insert
	} else if insert.key < trailing.key { // on which branch of the parent to put insert?
		trailing.left = insert
	} else {
		trailing.right = insert
	}
	t.size++
	t.insertFixup(insert)
}

func (t *rbTree) insertFixup(z *rbNode) {
	for z.parent.color == true { // while insert's parent color is red
		if z.parent == z.parent.parent.left {
			uncle := z.parent.parent.right
			if uncle.color == true { // uncle is red
				// property 4 is violated.
				// By coloring insert.parent & uncle black this is fixed.
				// now insert.parent.parent (grandparent) has to be colored red to maintain property 5
				//           B       z ->  R
				//         /  \          /  \
				//        R    R  ->    B    B
				//       /             /
				// z -> R             R
				z.parent.color = false       // case 1
				uncle.color = false          // case 1
				z.parent.parent.color = true // case 1
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// property 4 is violated
					// if z is the right child it gets rotated to be the left, which transforms it into case 3.
					//     B             B
					//   /  \          /  \
					//  R       ->    R
					//   \           /
					//    R        R
					// z has to get moved up because the rotation will move it down
					z = z.parent    // case 2
					t.leftRotate(z) // case 2
				}
				// property 4 is violated
				// colors are adjusted and a right rotation is performed to preserve property 5.
				//       B              B
				//     /  \           /  \
				//    R        ->    R    R
				//   /                     \
				//  R
				z.parent.color = false         // case 3
				z.parent.parent.color = true   // case 3
				t.rightRotate(z.parent.parent) // case 3
				// the loop terminates because z.parent is set to black
			}
		} else {
			// if z's parent is a left child the cases remain the same but left and
			// right are swapped.
			uncle := z.parent.parent.left
			if uncle.color {
				z.parent.color = false       // case 1
				uncle.color = false          // case 1
				z.parent.parent.color = true // case 1
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent     // case 2
					t.rightRotate(z) // case 2
				}
				z.parent.color = false        // case 3
				z.parent.parent.color = true  // case 3
				t.leftRotate(z.parent.parent) // case 3
			}
		}
	}
	t.root.color = false
}

func (t *rbTree) Delete(key int) {
	z := t.search(key)
	if z == nil {
		return
	}
	del, delOrigColor := z, z.color
	var move *rbNode
	// If z has only one child then it can just replaced by that one.
	if z.left == t.sentinel {
		move = z.right
		t.transplant(z, move)
	} else if z.right == t.sentinel {
		move = z.left
		t.transplant(z, move)
	} else {
		del = t.Min(z.right)
		delOrigColor = del.color
		move = del.right
		if del.parent == z {
			move.parent = del
		} else {
			t.transplant(del, move)
			del.right = z.right
			z.right.parent = del
		}
		t.transplant(z, del)
		del.left = z.left
		del.left.parent = del
		del.color = z.color
	}
	t.size--
	// if delOrigColor is black the deletion of node z can cause a violation of
	// property 5.
	if delOrigColor == false {
		t.deleteFixup(move)
	}
}

func (t *rbTree) transplant(u, v *rbNode) {
	if u.parent == t.sentinel {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *rbTree) deleteFixup(z *rbNode) {
	for z != t.root && z.color == false {
		if z == z.parent.left {
			w := z.parent.right
			if w.color == true {
				w.color = false        // case 1
				z.parent.color = true  // case 1
				t.leftRotate(z.parent) // case 1
				w = z.parent.right     // case 1
			}
			if w.left.color == false && w.right.color == false {
				w.color = true // case 2
				z = z.parent   // case 2
			} else {
				if w.right.color == false {
					w.left.color = false // case 3
					w.color = true       // case 3
					t.rightRotate(w)     // case 3
					w = z.parent.right   // case 3
				}
				w.color = z.parent.color // case 4
				z.parent.color = false   // case 4
				w.right.color = false    // case 4
				t.leftRotate(z.parent)   // case 4
				z = t.root               // case 4
			}
		} else {
			w := z.parent.left
			if w.color == true {
				w.color = false         // case 1
				z.parent.color = true   // case 1
				t.rightRotate(z.parent) // case 1
				w = z.parent.right      // case 1
			}
			if w.right.color == false && w.left.color == false {
				w.color = true // case 2
				z = z.parent   // case 2
			} else {
				if w.left.color == false {
					w.right.color = false // case 3
					w.color = true        // case 3
					t.leftRotate(w)       // case 3
					w = z.parent.left     // case 3
				}
				w.color = z.parent.color // case 4
				z.parent.color = false   // case 4
				w.left.color = false     // case 4
				t.rightRotate(z.parent)  // case 4
				z = t.root               // case 4
			}
		}
	}
	z.color = false
}

func (t *rbTree) Predecessor(node *rbNode) *rbNode {
	if node == nil {
		return nil
	}
	if node.left != t.sentinel {
		return t.Max(node.left)
	}
	parent := node.parent
	for parent != t.sentinel && node == parent.left {
		node = parent
		parent = parent.parent
	}
	if parent == t.sentinel {
		return nil
	}
	return parent
}

func (t *rbTree) Successor(node *rbNode) *rbNode {
	if node == nil {
		return nil
	}
	if node.right != t.sentinel {
		return t.Min(node.right)
	}
	parent := node.parent
	for parent != t.sentinel && node == parent.right {
		node = parent
		parent = parent.parent
	}
	if parent == t.sentinel {
		return nil
	}
	return parent
}

// Min finds the node with the smallest key in the tree.
func (t *rbTree) Min(n *rbNode) *rbNode {
	for n.left != t.sentinel {
		n = n.left
	}
	return n
}

// Max finds the node with the gratest key in the tree.
func (t *rbTree) Max(n *rbNode) *rbNode {
	for n.right != t.sentinel {
		n = n.right
	}
	return n
}

// leftRotate rotates the nodes connected to node around the axis from node to node.right.
// The binary search tree property is preserved by this operation.
// Ex:
//		x					y
//    /  \				  /  \
//   A    y 	->  	 x    C
//  	/  \		   /  \
//     B    C		  A    B
func (t *rbTree) leftRotate(node *rbNode) {
	if node == nil || node.right == nil { // nodes right child cannot be nil
		return
	}
	rot := node.right
	node.right = rot.left // make rot's left subtree node's right
	if rot.left != t.sentinel {
		rot.left.parent = node
	}
	rot.parent = node.parent       // link node's parent to rot's
	if node.parent == t.sentinel { // make replace node with rot in node's parent
		t.root = rot
	} else if node == node.parent.left {
		node.parent.left = rot
	} else {
		node.parent.right = rot
	}
	rot.left = node // put node on rot's left
	node.parent = rot
}

// rightRotate rotates the nodes connected to node around the axis from node to node.left
// The binary search tree property is preserved by this operation.
// Ex:
//  		y				x
//  	  /  \			  /  \
//  	 x    C		->	 A    y
//     /  \					/  \
//	  A    B			   B 	C
func (t *rbTree) rightRotate(node *rbNode) {
	if node == nil || node.left == nil { // nodes left child cannot be nil
		return
	}
	rot := node.left
	node.left = rot.right // make rot's right subtree node's left
	if rot.right != t.sentinel {
		rot.right.parent = node
	}
	rot.parent = node.parent       // link node's parent to rot's
	if node.parent == t.sentinel { // make replace node with rot in node's parent
		t.root = rot
	} else if node == node.parent.left {
		node.parent.left = rot
	} else {
		node.parent.right = rot
	}
	rot.right = node // put node on rot's right
	node.parent = rot
}

func (t *rbTree) String() string {
	return print(t.root, "", true)
}

func (n *rbNode) String() string {
	color := "BLACK"
	if n.color {
		color = "RED"
	}
	return fmt.Sprintf("%s k: %d,v: %v", color, n.key, n.value)
}

func print(node *rbNode, prefix string, tail bool) string {
	if node == nil {
		return ""
	}
	spacer1, spacer2 := "└── ", "    "
	if !tail {
		spacer1 = "├── "
		spacer2 = "│    "
	}
	str := helper.Concat(prefix, spacer1, node.String(), "\n")
	w := helper.Concat(prefix, spacer2)
	right := print(node.right, w, false)
	left := print(node.left, w, true)
	return helper.Concat(str, right, left)
}
