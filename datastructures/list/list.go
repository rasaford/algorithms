package list

import (
	"fmt"
	"math"
)

type Node struct {
	value      int
	list       *List
	prev, next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	sentinel *Node
	len      int
}

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
func (l *List) Start() *Node {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.next
}
func (l *List) End() *Node {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.prev
}

func (l *List) Len() int {
	return l.len
}

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

func (l *List) Delete(node *Node) error {
	if !l.validNode(node) {
		return fmt.Errorf("cannot delete node:%v from list", node)
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.len--
	return nil
}

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

func (l *List) validNode(node *Node) bool {
	if node == nil || node.prev == nil ||
		node.next == nil || node.list != l ||
		node == l.sentinel {
		return false
	}
	return true
}

// func (l *List) InsertAfter(val int, node *Node) {
// 	next := node.next
// 	n := &Node{
// 		value: val,
// 		list:  l,
// 		prev:  node,
// 		next:  next,
// 	}
// 	node.next = n
// 4}

// func (l *List) InsertBefore(val int, node *Node) {
// 	prev := node.prev
// 	n := &Node{
// 		value: val,
// 		list:  l,
// 		prev:  prev,
// 		next:  node,
// 	}
// 	node.prev = n
// }
