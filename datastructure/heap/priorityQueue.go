package heap

import (
	"fmt"
	"math"
)

// PriorityQueue is a datastructure that sortes all inserted keys in sorted order and
// allows for O(1) access time for the Head of the queue.
// All other operations take O(lg n) n:= size(queue) time
type PriorityQueue interface {
	Insert(int) error
	Head() int
	ExtractHead() (int, error)
	UpdateKey(int, int) error
}

type queue struct {
	*Heap
}

// NewMaxPriorityQueue builds an empty priority queue basend on a max-heap.
// The element with the highest priority is always at the root of the tree
// and can be accessed in O(1) time.
func NewMaxPriorityQueue(array []int) PriorityQueue {
	return &queue{NewMaxHeap(array)}
}

// NewMinPriorityQueue builds an empty priority queue basend on a min-heap.
// The element with the lowest priority is always at the root of the tree
// and can be accessed in O(1) time.
func NewMinPriorityQueue(array []int) PriorityQueue {
	return &queue{NewMinHeap(array)}
}

// Insert inserts a new element into the queue at the index corresponding
// to the priority of the given key.
//
// If the new key cannot be inserted because it is to small / large for the used
// heap type an error is thrown.
func (q *queue) Insert(key int) error {
	newTemp := math.MinInt32
	if q.hType == "min-heap" {
		newTemp = math.MaxInt32
	}
	q.heap = append(q.heap, newTemp)
	q.size++
	return q.UpdateKey(q.size, key)
}

// Head returns the highest / lowest priority element
// depending on the queue type.
func (q *queue) Head() int {
	return q.heap[0]
}

// ExtractHead returns the highest / lowest priority element and delets it
// depending on the queue type.
func (q *queue) ExtractHead() (int, error) {
	if q.size < 0 {
		return -1, fmt.Errorf("the heap is too small (size %d)", q.size)
	}
	max := q.heap[0]
	q.heap[0] = q.heap[q.size]
	q.size--
	q.Heapify(0)
	return max, nil
}

// UpdateKey updates the key value at the specified index to the new key.
// The new key value has to be larger / smaller than the current one depending on the
// heap type.
// If this is not the case an error is thrown.
func (q *queue) UpdateKey(index, key int) error {
	if !q.compare(key, q.heap[index]) {
		errType := "smaller"
		if q.hType == "min-heap" {
			errType = "larger"
		}
		return fmt.Errorf("the new key is %s than the current one", errType)
	}
	q.heap[index] = key
	for index > 0 && q.compare(q.heap[index], q.heap[parent(index)]) {
		swap(&q.heap[index], &q.heap[parent(index)])
		index = parent(index)
	}
	return nil
}
