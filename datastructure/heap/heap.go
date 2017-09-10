package heap

import "math"

type Heap struct {
	heap    []int
	size    int
	compare func(int, int) bool
}

// NewMaxHeap builds a max Heap structure on the given array. Every node in the
// built binary tree is larger than it's two children.
// The tree is represented by the array.
// Example:
//    5
//  /  \
// 3   1
func NewMaxHeap(array []int) Heap {
	heap := Heap{
		heap: array,
		size: len(array) - 1,
		compare: func(a, b int) bool {
			if a > b {
				return true
			}
			return false
		},
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
func NewMinHeap(array []int) Heap {
	heap := Heap{
		heap: array,
		size: len(array) - 1,
		compare: func(a, b int) bool {
			if a < b {
				return true
			}
			return false
		},
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
// It can be calculated by parent := cleil(index/2) - 1
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
		swap(&h.heap[index], &h.heap[largest])
		h.Heapify(largest)
	}
}

// heapifyIter is a non-recursive implementation of the heapfiy procedure.
func (h *Heap) heapifyIter(index int) {
	largest := -1
	temporaryIndex := index
	lChild, rChild := 0, 0
	for largest != index {
		largest, index = temporaryIndex, temporaryIndex
		lChild = left(index)
		rChild = right(index)
		if lChild <= h.size && h.compare(h.heap[lChild], h.heap[index]) {
			largest = lChild
		}
		if rChild <= h.size && h.compare(h.heap[rChild], h.heap[largest]) {
			largest = rChild
		}
		swap(&h.heap[index], &h.heap[largest])
		temporaryIndex = largest
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
