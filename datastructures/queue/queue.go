package queue

import (
	"fmt"
)

// Queue is an implementation of a FIFO datastructure.
type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
}

type boundedQueue struct {
	array      []interface{}
	tail, head int
}

// NewBounded returns a new, empty queue with a fixed size.
func NewBounded(size int) Queue {
	return &boundedQueue{
		array: make([]interface{}, size+1),
		tail:  1,
		head:  1,
	}
}

// Enqueue adds a new Element at the beginning of the queue.
//
// It runs in O(1) time.
func (q *boundedQueue) Enqueue(value interface{}) error {
	if q.head == q.tail+1 {
		return fmt.Errorf("cannot enqueue to a full queue")
	}
	q.array[q.tail] = value
	q.tail = next(q.tail, len(q.array))
	return nil
}

// Dequeue removes the first element from the queue.
//
// It runs in O(1) time.
func (q *boundedQueue) Dequeue() (interface{}, error) {
	if q.head == q.tail {
		return nil, fmt.Errorf("cannot dequeue from an empty queue")
	}
	ret := q.array[q.head]
	q.head = next(q.head, len(q.array))
	return ret, nil
}

type unboundedQueue struct {
	array []interface{}
}

// NewUnbounded returns a new, empty queue with "infinite" size.
func NewUnbounded() Queue {
	return &unboundedQueue{
		array: make([]interface{}, 0, 8),
	}
}

// Enqueue adds a new Element at the beginning of the queue.
//
// It runs in O(1) time.
func (q *unboundedQueue) Enqueue(value interface{}) error {
	q.array = append(q.array, value)
	return nil
}

// Dequeue removes the first element from the queue.
//
// It runs in O(1) time.
func (q *unboundedQueue) Dequeue() (interface{}, error) {
	if len(q.array) == 0 {
		return nil, fmt.Errorf("cannot dequeue from an empty queue")
	}
	ret := q.array[0]
	q.array = q.array[1:]
	return ret, nil
}

// next calculates the index of the next element in a looping array.
func next(v int, max int) int {
	return (v + 1) % max
}

// prev calculates the index of the previous element in a looping array.
func prev(v int, max int) int {
	res := (v - 1) % max
	if res < 0 {
		res = max + res
	}
	return res
}
