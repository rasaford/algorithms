package list

import (
	"fmt"
	"math"
	"reflect"
)

// Node represents one node of the list with links to the
// previous and next node, as well as the list.
type Node struct {
	// value stored in the node
	Value interface{}
	list  *List
	// pointers to the previous and next node in the list
	Prev, Next *Node
}

// List is the container for a list
type List struct {
	sentinel *Node
	// Len is the length of the list
	Len int
}

// New creates a circular doubly linked list with a sentinel value.
func New() *List {
	list := &List{}
	// the first node in a list is always a sentinel item with no value.
	sentinel := &Node{
		list:  list,
		Value: math.MinInt32,
	}
	sentinel.Next = sentinel
	sentinel.Prev = sentinel
	list.sentinel = sentinel
	return list
}

// Start returns a pointer to the first node in the list
func (l *List) Start() *Node {
	if l.Len == 0 {
		return nil
	}
	return l.sentinel.Next
}

// End returns a pointer to the last node in the list
func (l *List) End() *Node {
	if l.Len == 0 {
		return nil
	}
	return l.sentinel.Prev
}

// Insert appends a new value to the list.
//
// It runs ins O(1)
func (l *List) Insert(val interface{}) {
	prev := l.sentinel.Prev
	node := &Node{
		Value: val,
		list:  l,
		Prev:  prev,
		Next:  l.sentinel,
	}
	l.sentinel.Prev = node
	prev.Next = node
	l.Len++
}

// Delete deletes the given node from the list.
// If the given node is invalid an error is returned.
func (l *List) Delete(node *Node) error {
	if !l.validNode(node) {
		return fmt.Errorf("cannot delete node:%v from list", node)
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	l.Len--
	return nil
}

// Search searches for the key in the list.
// If the nodes cannot be found an error is returned.
//
// It runs in O(n) time with n := len(list)
func (l *List) Search(key interface{}) (*Node, error) {
	element := l.sentinel.Next
	for element != l.sentinel && !reflect.DeepEqual(element.Value, key) {
		element = element.Next
	}
	if element == l.sentinel {
		return nil, fmt.Errorf("key: %d is not contained in the list", key)
	}
	return element, nil
}

// checks if the given node is in the list.
func (l *List) validNode(node *Node) bool {
	if node == nil || node.Prev == nil ||
		node.Next == nil || node.list != l ||
		node == l.sentinel {
		return false
	}
	return true
}

// InsertAfter inserts the value after the given node.
//
// It runs ins O(1)
func (l *List) InsertAfter(val interface{}, node *Node) {
	if node == nil {
		return
	}
	next := node.Next
	n := &Node{
		Value: val,
		list:  l,
		Prev:  node,
		Next:  next,
	}
	node.Next = n
}

// InsertBefore inserts the value before the given node.
//
// It runs ins O(1)
func (l *List) InsertBefore(val interface{}, node *Node) {
	if node == nil {
		return
	}
	prev := node.Prev
	n := &Node{
		Value: val,
		list:  l,
		Prev:  prev,
		Next:  node,
	}
	node.Prev = n
}

// Concat concatenates two lists and retuns the result
//
// It runs in O(1) time
func Concat(a, b *List) *List {
	if a.Len == 0 && b.Len == 0 {
		return a
	} else if a.Len == 0 {
		return b
	} else if b.Len == 0 {
		return a
	}
	aEnd := a.End()
	bStart := b.Start()
	bEnd := b.End()
	aEnd.Next = bStart
	bStart.Prev = aEnd
	a.sentinel.Prev = bEnd
	bEnd.Next = a.sentinel
	a.Len += b.Len
	return a
}
