package heap

import (
	"fmt"
	"math"

	"github.com/rasaford/algorithms/internal/helper"
)

// Heap contains the underlying array of the Heap tree structure. It also stores
// the comparator as well as the current size of the Heap.
type Heap struct {
	heap []int
	size int
	// compare defines weather the first element is larger / smaller than
	// the second based on the type of heap used.
	compare func(int, int) bool
	hType   string
}

// NewMaxHeap builds a max Heap structure on the given array. Every node in the
// built binary tree is larger than it's two children.
// The tree is represented by the array.
// Example:
//    5
//  /  \
// 3   1
func NewMaxHeap(array []int) *Heap {
	heap := &Heap{
		heap: array,
		size: len(array) - 1,
		compare: func(a, b int) bool {
			if a > b {
				return true
			}
			return false
		},
		hType: "max-heap",
	}
	for i := heap.size / 2; i >= 0; i-- {
		heap.Heapify(i)
	}
	return heap
}

// NewMinHeap builds a min Heap structure on the given array. Every node in the
// built binary tree is smaller than it's two children.
// The tree is represented by the array.
// Example:
//    1
//  /  \
// 2   5
func NewMinHeap(array []int) *Heap {
	heap := &Heap{
		heap: array,
		size: len(array) - 1,
		compare: func(a, b int) bool {
			if a < b {
				return true
			}
			return false
		},
		hType: "min-heap",
	}
	for i := heap.size / 2; i >= 0; i-- {
		heap.Heapify(i)
	}
	return heap
}

// Size returns the current number of elements in the Heap structure.
func (h *Heap) Size() int {
	return h.size
}

// Decrement decrements the size of the Heap structure by 1 therefore
// making the last element of the underlying array invisible to the Heapify procedure.
// The underlying array however is not changed
func (h *Heap) Decrement() {
	h.size--
}

// parent returns the index of the parent node in the heap.
//
// It can be calculated by parent := ceil(index/2) - 1
func parent(index int) int {
	return int(math.Ceil(float64(index)/2)) - 1
}

// left returns the index of the left child node in the heap.
//
// It can be calculated by left := floor(index*2) + 1
func left(index int) int {
	return index<<1 + 1
}

// right returns the index of the right child node in the heap.
//
// It can be calculated by right:= floor(index*2) + 2
func right(index int) int {
	return index<<1 + 2 // equivalent to index * 2 + 2
}

// Heapify assumes the element at the given index is the only element in the tree
// that violates the heap property. It recursively swaps the element at index to the right
// place
// Runtime: O(lg n)
func (h *Heap) Heapify(index int) {
	lChild := left(index)
	rChild := right(index)
	largest := index
	if lChild <= h.size && h.compare(h.heap[lChild], h.heap[index]) {
		largest = lChild
	}
	if rChild <= h.size && h.compare(h.heap[rChild], h.heap[largest]) {
		largest = rChild
	}
	if largest != index {
		helper.Swap(&h.heap[index], &h.heap[largest])
		h.Heapify(largest)
	}
}

// heapifyIter is a non-recursive implementation of the heapfiy procedure.
// func (h *Heap) heapifyIter(index int) {
// 	largest := -1
// 	temporaryIndex := index
// 	lChild, rChild := 0, 0
// 	for largest != index {
// 		largest, index = temporaryIndex, temporaryIndex
// 		lChild = left(index)
// 		rChild = right(index)
// 		if lChild <= h.size && h.compare(h.heap[lChild], h.heap[index]) {
// 			largest = lChild
// 		}
// 		if rChild <= h.size && h.compare(h.heap[rChild], h.heap[largest]) {
// 			largest = rChild
// 		}
// 		swap(&h.heap[index], &h.heap[largest])
// 		temporaryIndex = largest
// 	}
// }

// Insert inserts a new element into the queue at the index corresponding
// to the priority of the given key.
//
// If the new key cannot be inserted because it is to small / large for the used
// heap type an error is thrown.
func (h *Heap) Insert(key int) error {
	newTemp := math.MinInt32
	if h.hType == "min-heap" {
		newTemp = math.MaxInt32
	}
	h.heap = append(h.heap, newTemp)
	h.size++
	return h.UpdateKey(h.size, key)
}

// Head returns the highest / lowest priority element
// depending on the queue type.
func (h *Heap) Head() int {
	return h.heap[0]
}

// ExtractHead returns the highest / lowest priority element and deletes it
// depending on the queue type.
func (h *Heap) ExtractHead() (int, error) {
	if h.size < 0 {
		return -1, fmt.Errorf("the heap is too small (size %d)", h.size)
	}
	max := h.heap[0]
	h.heap[0] = h.heap[h.size]
	h.size--
	h.Heapify(0)
	return max, nil
}

// UpdateKey updates the key value at the specified index to the new key.
// The new key value has to be larger / smaller than the current one depending on the
// heap type.
// If this is not the case an error is thrown.
func (h *Heap) UpdateKey(index, key int) error {
	if !h.compare(key, h.heap[index]) {
		errType := "smaller"
		if h.hType == "min-heap" {
			errType = "larger"
		}
		return fmt.Errorf("the new key is %s than the current one", errType)
	}
	for index > 0 && h.compare(key, h.heap[parent(index)]) {
		h.heap[index] = h.heap[parent(index)]
		index = parent(index)
	}
	h.heap[index] = key
	return nil
}
