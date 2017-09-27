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
	root     RedBlackNode
	sentinel RedBlackNode
	size     int
}

type RedBlackNode interface {
	Key() int
	setKey(int)
	Value() interface{}
	setValue(interface{})
	Left() RedBlackNode
	setLeft(RedBlackNode)
	Right() RedBlackNode
	setRight(RedBlackNode)
	Parent() RedBlackNode
	setParent(RedBlackNode)
	Color() bool
	setColor(bool)
	String() string
}

type rbNode struct {
	key                 int
	value               interface{}
	left, right, parent RedBlackNode
	// color of the node in the RB-Tree. true is red, false is black
	color bool
}

func (n *rbNode) Key() int                    { return n.key }
func (n *rbNode) setKey(k int)                { n.key = k }
func (n *rbNode) Value() interface{}          { return n.value }
func (n *rbNode) setValue(v interface{})      { n.value = v }
func (n *rbNode) Left() RedBlackNode          { return n.left }
func (n *rbNode) setLeft(node RedBlackNode)   { n.left = node }
func (n *rbNode) Right() RedBlackNode         { return n.right }
func (n *rbNode) setRight(node RedBlackNode)  { n.right = node }
func (n *rbNode) Parent() RedBlackNode        { return n.parent }
func (n *rbNode) setParent(node RedBlackNode) { n.parent = node }
func (n *rbNode) Color() bool                 { return n.color }
func (n *rbNode) setColor(col bool)           { n.color = col }

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
		return node.Value()
	}
	return nil
}

func (t *rbTree) search(key int) RedBlackNode {
	n := t.root
	for n != nil && key != n.Key() {
		if key < n.Key() {
			n = n.Left()
		} else {
			n = n.Right()
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
	t.insert(insert)
}

func (t *rbTree) insert(insert RedBlackNode) {
	// maintains a training pointer to get back to the parent.
	trailing, node := t.sentinel, t.root
	for node != t.sentinel {
		trailing = node
		if insert.Key() < node.Key() {
			node = node.Left()
		} else {
			node = node.Right()
		}
	}
	insert.setParent(trailing)
	if trailing == t.sentinel {
		t.root = insert
	} else if insert.Key() < trailing.Key() { // on which branch of the parent to put insert?
		trailing.setLeft(insert)
	} else {
		trailing.setRight(insert)
	}
	t.size++
	t.insertFixup(insert)
}

func (t *rbTree) insertFixup(z RedBlackNode) {
	for z.Parent().Color() == true { // while insert's parent color is red
		if z.Parent() == z.Parent().Parent().Left() {
			uncle := z.Parent().Parent().Right()
			if uncle.Color() == true { // uncle is red
				// property 4 is violated.
				// By coloring insert.parent & uncle black this is fixed.
				// now insert.parent.parent (grandparent) has to be colored red to maintain property 5
				//           B       z ->  R
				//         /  \          /  \
				//        R    R  ->    B    B
				//       /             /
				// z -> R             R
				z.Parent().setColor(false)         // case 1
				uncle.setColor(false)              // case 1
				z.Parent().Parent().setColor(true) // case 1
				z = z.Parent().Parent()
			} else {
				if z == z.Parent().Right() {
					// property 4 is violated
					// if z is the right child it gets rotated to be the left, which transforms it into case 3.
					//     B             B
					//   /  \          /  \
					//  R       ->    R
					//   \           /
					//    R        R
					// z has to get moved up because the rotation will move it down
					z = z.Parent()  // case 2
					t.leftRotate(z) // case 2
				}
				// property 4 is violated
				// colors are adjusted and a right rotation is performed to preserve property 5.
				//       B              B
				//     /  \           /  \
				//    R        ->    R    R
				//   /                     \
				//  R
				z.Parent().setColor(false)         // case 3
				z.Parent().Parent().setColor(true) // case 3
				t.rightRotate(z.Parent().Parent()) // case 3
				// the loop terminates because z.parent is set to black
			}
		} else {
			// if z's parent is a left child the cases remain the same but left and
			// right are swapped.
			uncle := z.Parent().Parent().Left()
			if uncle.Color() == true {
				z.Parent().setColor(false)         // case 1
				uncle.setColor(false)              // case 1
				z.Parent().Parent().setColor(true) // case 1
				z = z.Parent().Parent()
			} else {
				if z == z.Parent().Left() {
					z = z.Parent()   // case 2
					t.rightRotate(z) // case 2
				}
				z.Parent().setColor(false)         // case 3
				z.Parent().Parent().setColor(true) // case 3
				t.leftRotate(z.Parent().Parent())  // case 3
			}
		}
	}
	t.root.setColor(false)
}

func (t *rbTree) Delete(key int) {
	z := t.search(key)
	if z == nil {
		return
	}
	t.delete(z)
}

func (t *rbTree) delete(z RedBlackNode) {
	y, yOrig := z, z.Color()
	var x RedBlackNode
	// If z has only one child then it can just replaced by that one.
	if z.Left() == t.sentinel {
		x = z.Right()
		t.transplant(z, z.Right())
	} else if z.Right() == t.sentinel {
		x = z.Left()
		t.transplant(z, z.Left())
	} else {
		y = t.Min(z.Right())
		yOrig = y.Color()
		x = y.Right()
		if y.Parent() == z {
			x.setParent(y)
		} else {
			t.transplant(y, y.Right())
			y.setRight(z.Right())
			z.Right().setParent(y)
		}
		t.transplant(z, y)
		y.setLeft(z.Left())
		y.Left().setParent(y)
		y.setColor(z.Color())
	}
	t.size--
	// if delOrigColor is black the deletion of node z can cause a violation of
	// property 5.
	if yOrig == false {
		t.deleteFixup(x)
	}
}

func (t *rbTree) transplant(u, v RedBlackNode) {
	if u.Parent() == t.sentinel {
		t.root = v
	} else if u == u.Parent().Left() {
		u.Parent().setLeft(v)
	} else {
		u.Parent().setRight(v)
	}
	v.setParent(u.Parent())
}

func (t *rbTree) deleteFixup(x RedBlackNode) {
	for x != t.root && x.Color() == false {
		if x == x.Parent().Left() {
			w := x.Parent().Right()
			if w.Color() == true {
				w.setColor(false)         // case 1
				x.Parent().setColor(true) // case 1
				t.leftRotate(x.Parent())  // case 1
				w = x.Parent().Right()    // case 1
			}
			if w.Left().Color() == false && w.Right().Color() == false {
				w.setColor(true) // case 2
				x = x.Parent()   // case 2
			} else {
				if w.Right().Color() == false {
					w.Left().setColor(false) // case 3
					w.setColor(true)         // case 3
					t.rightRotate(w)         // case 3
					w = x.Parent().Right()   // case 3
				}
				w.setColor(x.Parent().Color()) // case 4
				x.Parent().setColor(false)     // case 4
				w.Right().setColor(false)      // case 4
				t.leftRotate(x.Parent())       // case 4
				x = t.root                     // case 4
			}
		} else {
			w := x.Parent().Left()
			if w.Color() == true {
				w.setColor(false)         // case 1
				x.Parent().setColor(true) // case 1
				t.rightRotate(x.Parent()) // case 1
				w = x.Parent().Right()    // case 1
			}
			if w.Right().Color() == false && w.Left().Color() == false {
				w.setColor(true) // case 2
				x = x.Parent()   // case 2
			} else {
				if w.Left().Color() == false {
					w.Right().setColor(false) // case 3
					w.setColor(true)          // case 3
					t.leftRotate(w)           // case 3
					w = x.Parent().Left()     // case 3
				}
				w.setColor(x.Parent().Color()) // case 4
				x.Parent().setColor(false)     // case 4
				w.Left().setColor(false)       // case 4
				t.rightRotate(x.Parent())      // case 4
				x = t.root                     // case 4
			}
		}
	}
	x.setColor(false)
}

func (t *rbTree) Predecessor(node RedBlackNode) RedBlackNode {
	if node == nil {
		return nil
	}
	if node.Left() != t.sentinel {
		return t.Max(node.Left())
	}
	parent := node.Parent()
	for parent != t.sentinel && node == parent.Left() {
		node = parent
		parent = parent.Parent()
	}
	if parent == t.sentinel {
		return nil
	}
	return parent
}

func (t *rbTree) Successor(node RedBlackNode) RedBlackNode {
	if node == nil {
		return nil
	}
	if node.Right() != t.sentinel {
		return t.Min(node.Right())
	}
	parent := node.Parent()
	for parent != t.sentinel && node == parent.Right() {
		node = parent
		parent = parent.Parent()
	}
	if parent == t.sentinel {
		return nil
	}
	return parent
}

// Min finds the node with the smallest key in the tree.
func (t *rbTree) Min(n RedBlackNode) RedBlackNode {
	for n.Left() != t.sentinel {
		n = n.Left()
	}
	return n
}

// Max finds the node with the gratest key in the tree.
func (t *rbTree) Max(n RedBlackNode) RedBlackNode {
	for n.Right() != t.sentinel {
		n = n.Right()
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
func (t *rbTree) leftRotate(node RedBlackNode) {
	if node == nil || node.Right() == nil { // nodes right child cannot be nil
		return
	}
	rot := node.Right()
	node.setRight(rot.Left()) // make rot's left subtree node's right
	if rot.Left() != t.sentinel {
		rot.Left().setParent(node)
	}
	rot.setParent(node.Parent())     // link node's parent to rot's
	if node.Parent() == t.sentinel { // make replace node with rot in node's parent
		t.root = rot
	} else if node == node.Parent().Left() {
		node.Parent().setLeft(rot)
	} else {
		node.Parent().setRight(rot)
	}
	rot.setLeft(node) // put node on rot's left
	node.setParent(rot)
}

// rightRotate rotates the nodes connected to node around the axis from node to node.left
// The binary search tree property is preserved by this operation.
// Ex:
//  		y				x
//  	  /  \			  /  \
//  	 x    C		->	 A    y
//     /  \					/  \
//	  A    B			   B 	C
func (t *rbTree) rightRotate(node RedBlackNode) {
	if node == nil || node.Left() == nil { // nodes left child cannot be nil
		return
	}
	rot := node.Left()
	node.setLeft(rot.Right()) // make rot's right subtree node's left
	if rot.Right() != t.sentinel {
		rot.Right().setParent(node)
	}
	rot.setParent(node.Parent())     // link node's parent to rot's
	if node.Parent() == t.sentinel { // make replace node with rot in node's parent
		t.root = rot
	} else if node == node.Parent().Left() {
		node.Parent().setLeft(rot)
	} else {
		node.Parent().setRight(rot)
	}
	rot.setRight(node) // put node on rot's right
	node.setParent(rot)
}

func (t *rbTree) String() string {
	return print(t.sentinel, t.root, "", true)
}

func (n *rbNode) String() string {
	color := "BLACK"
	if n.color {
		color = "RED"
	}
	return fmt.Sprintf("%5s k: %d,v: %v", color, n.key, n.value)
}

func print(sentinel, node RedBlackNode, prefix string, tail bool) string {
	if node == sentinel {
		return ""
	}
	spacer1, spacer2 := "└── ", "    "
	if !tail {
		spacer1 = "├── "
		spacer2 = "│    "
	}
	str := helper.Concat(prefix, spacer1, node.String(), "\n")
	w := helper.Concat(prefix, spacer2)
	right := print(sentinel, node.Right(), w, false)
	left := print(sentinel, node.Left(), w, true)
	return helper.Concat(str, right, left)
}
