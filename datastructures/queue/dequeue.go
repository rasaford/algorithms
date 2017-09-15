package queue

import (
	"fmt"
	"math"
)

// Dequeue is an implementation of a FIFO datastructure.
// It allows for adding and removing elements from both the head & tail of
// the queue.
type Dequeue interface {
	EnqueueTail(interface{}) error
	DequeueTail() (interface{}, error)
	EnqueueHead(interface{}) error
	DequeueHead() (interface{}, error)
}

type deQueue struct {
	array      []interface{}
	head, tail int
}

// NewDequeue returns a new, empty dequeue with a fixed size.
func NewDequeue(size int) Dequeue {
	size = int(math.Max(float64(size), 2))
	return &deQueue{
		array: make([]interface{}, size+1),
		head:  size / 2,
		tail:  size / 2,
	}
}

// EnqueueHead adds a new Element at the beginning of the queue.
//
// It runs in O(1) time.
func (q *deQueue) EnqueueHead(value interface{}) error {
	if q.tail == prev(q.head, len(q.array)) {
		return fmt.Errorf("cannot enqueue to a full queue")
	}
	q.array[q.head] = value
	q.head = prev(q.head, len(q.array))
	return nil
}

// EnqueueTail adds a new Element at the end od the queue.
//
// It runs in O(1) time.
func (q *deQueue) EnqueueTail(value interface{}) error {
	if q.head == next(q.tail, len(q.array)) {
		return fmt.Errorf("cannot enqueue to a full queue")
	}
	q.array[q.tail] = value
	q.tail = next(q.tail, len(q.array))
	return nil
}

// DequeueHead removes the last element from the queue.
//
// It runs in O(1) time.
func (q *deQueue) DequeueHead() (interface{}, error) {
	if q.head == q.tail {
		return nil, fmt.Errorf("cannot dequeue from an empty queue")
	}
	ret := q.array[q.head]
	q.head = next(q.head, len(q.array))
	return ret, nil
}

// DequeueTail removes the first element from the queue.
//
// It runs in O(1) time.
func (q *deQueue) DequeueTail() (interface{}, error) {
	if q.head == q.tail {
		return nil, fmt.Errorf("cannot dequeue from an empty queue")
	}
	ret := q.array[q.tail]
	q.tail = prev(q.tail, len(q.array))
	return ret, nil
}
