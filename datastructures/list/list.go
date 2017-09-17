package list

import (
	"fmt"
	"math"
)

// Node represents one node of the list with links to the
// previous and next node, as well as the list.
type Node struct {
	value      int
	list       *List
	prev, next *Node
}

// Next returns a pointer to the next node in the list
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node in the list
func (n *Node) Prev() *Node {
	return n.prev
}

// List is the container for a list
type List struct {
	sentinel *Node
	len      int
}

// New creates a circular doubly linked list with a sentinel value.
func New() *List {
	list := &List{}
	// the first node in a list is always a sentinel item with no value.
	sentinel := &Node{
		list:  list,
		value: math.MinInt32,
	}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	list.sentinel = sentinel
	return list
}

// Start returns a pointer to the first node in the list
func (l *List) Start() *Node {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.next
}

// End returns a pointer to the last node in the list
func (l *List) End() *Node {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.prev
}

// Len returns the length of the list
func (l *List) Len() int {
	return l.len
}

// Insert appends a new value to the list.
//
// It runs ins O(1)
func (l *List) Insert(val int) {
	prev := l.sentinel.prev
	node := &Node{
		value: val,
		list:  l,
		prev:  prev,
		next:  l.sentinel,
	}
	l.sentinel.prev = node
	prev.next = node
	l.len++
}

// Delete deletes the given node from the list.
// If the given node is invalid an error is returned.
func (l *List) Delete(node *Node) error {
	if !l.validNode(node) {
		return fmt.Errorf("cannot delete node:%v from list", node)
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.len--
	return nil
}

// Search searches for the key in the list.
// If the nodes cannot be found an error is returned.
//
// It runs in O(n) time with n := len(list)
func (l *List) Search(key int) (*Node, error) {
	element := l.sentinel.next
	for element != l.sentinel && element.value != key {
		element = element.next
	}
	if element == l.sentinel {
		return nil, fmt.Errorf("key: %d is not contained in the list", key)
	}
	return element, nil
}

// checks if the given node is in the list.
func (l *List) validNode(node *Node) bool {
	if node == nil || node.prev == nil ||
		node.next == nil || node.list != l ||
		node == l.sentinel {
		return false
	}
	return true
}

// InsertAfter inserts the value after the given node.
//
// It runs ins O(1)
func (l *List) InsertAfter(val int, node *Node) {
	if node == nil {
		return
	}
	next := node.next
	n := &Node{
		value: val,
		list:  l,
		prev:  node,
		next:  next,
	}
	node.next = n
}

// InsertBefore inserts the value before the given node.
//
// It runs ins O(1)
func (l *List) InsertBefore(val int, node *Node) {
	if node == nil {
		return
	}
	prev := node.prev
	n := &Node{
		value: val,
		list:  l,
		prev:  prev,
		next:  node,
	}
	node.prev = n
}

// Concat concatenates two lists and retuns the result
//
// It runs in O(1) time
func Concat(a, b *List) *List {
	if a.len == 0 && b.len == 0 {
		return a
	} else if a.len == 0 {
		return b
	} else if b.len == 0 {
		return a
	}
	aEnd := a.End()
	bStart := b.Start()
	bEnd := b.End()
	aEnd.next = bStart
	bStart.prev = aEnd
	a.sentinel.prev = bEnd
	bEnd.next = a.sentinel
	a.len += b.len
	return a
}
