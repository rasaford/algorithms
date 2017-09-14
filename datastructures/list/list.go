package list

import (
	"fmt"
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
	start, end *Node
	len        int
}

func New() *List {
	return &List{
		start: nil,
		end:   nil,
		len:   0,
	}
}
func (l *List) Start() *Node {
	return l.start
}
func (l *List) End() *Node {
	return l.end
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Remove(index int) error {
	node, err := l.Get(index)
	if err != nil {
		return err
	}
	prev, next := node.Prev(), node.Next()
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	return nil
}

func (l *List) Get(key int) (*Node, error) {
	node := l.start
	index := 0
	for node != nil {
		if index == key {
			return node, nil
		}
		node = node.next
		index++
	}
	return nil, fmt.Errorf("node with key: %d is not is the list", key)
}

func (l *List) InsertAfter(val int, node *Node) {
	next := node.next
	n := &Node{
		value: val,
		list:  l,
		prev:  node,
		next:  next,
	}
	node.next = n
	next.prev = n
	if l.start == nil {
		l.start = node
	}
	if l.end == nil {
		l.end = node
	}
}
