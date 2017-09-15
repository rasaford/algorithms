package sorting

import (
	"github.com/rasaford/algorithms/datastructures/heap"
	h "github.com/rasaford/algorithms/internal/helper"
)

// BubbleSort is the naive implementation of sorting an array.
// It is easy to implement but not very efficient.
//
// It runs in O(n^2) time, where n := len(input)
// Space Complexity is O(1)
func BubbleSort(input []int) []int {
	array := h.Clone(input)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				h.Swap(&array[i], &array[j])
			}
		}
	}
	return array
}

// InsertionSort is a simple sorting algorithm that scans
// through the array and maintains a sorted version of it
// behind the read head.
//
// It runs in O(n^2) time, where n := len(input)
// Space Complexity is O(1)
func InsertionSort(input []int) []int {
	array := h.Clone(input)
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}

// MergeSort is a more complex implmentation of sorting, it
// recursively splits the input and sorts the subarrays. Then
// it combines the inputs to form the original array. It therefore uses the
// divide and conquer paradigm.
//
// It runs in O(n lg n) time, where n := len(input)
// Space Complexity is O(n)
func MergeSort(input []int) []int {
	// merge sort needs 2 arrays
	// they get swaped at each level of the recursion tree
	array, work := h.Clone(input), h.Clone(input)
	mergeSortRec(array, work, 0, len(work))
	return array
}

// recursively splits the two arrays and swaps work and input at
// each step of the recursion tree
func mergeSortRec(input, work []int, low, high int) {
	if high-low <= 1 {
		return
	}
	mid := (low + high) / 2
	mergeSortRec(work, input, low, mid)
	mergeSortRec(work, input, mid, high)
	merge(work, input, low, mid, high)
}

// merges the two sorted subarrays from input to work
func merge(input, work []int, low, mid, high int) {
	i, j := low, mid
	for k := low; k < high; k++ {
		if i < mid && (j >= high || input[i] <= input[j]) {
			work[k] = input[i]
			i++
		} else {
			work[k] = input[j]
			j++
		}
	}
}

// HeapSort builds a max-Heap from the array and then iteratively extracts the
// maximum element.
// Every time an element is extracted the heap size gets decremented, effectively
// removing the element form the heap. The removed elements get stored in this
// unmodified part of the array.
//
// It runs in O(n lg n) time, where n := len(input)
// Space Complexity is O(1)
func HeapSort(input []int) []int {
	array := h.Clone(input)
	heap := heap.NewMaxHeap(array)
	for i := heap.Size(); i >= 1; i-- {
		h.Swap(&array[i], &array[0])
		heap.Decrement()
		heap.Heapify(0)
	}
	return array
}

// QuickSort sorts the input array by randomly selecting a radix point from a
// given subarray. Then it partitions that subarray into 2 sets, one with all elements
// that are <= the other one with all > than the radix point.
// This procedure is applied recursively to both partitions to sort the array.
//
// It runs in O(n lg n) time where n := len(input)
// Space Complexity is O(1)
func QuickSort(input []int) []int {
	array := h.Clone(input)
	quickSortRec(array, 0, len(array)-1)
	return array
}

func quickSortRec(input []int, start, end int) {
	if start >= end {
		return
	}
	mid := randomPartition(input, start, end)
	quickSortRec(input, start, mid-1)
	quickSortRec(input, mid+1, end)
}

func randomPartition(array []int, start, end int) int {
	rand, err := h.RandBetween(start, end)
	if err != nil {
		panic(err)
	}
	h.Swap(&array[rand], &array[end])
	radix := array[end]
	lowerBound := start - 1
	for i := start; i < end; i++ {
		if array[i] <= radix {
			lowerBound++
			h.Swap(&array[i], &array[lowerBound])
		}
	}
	lowerBound++
	h.Swap(&array[lowerBound], &array[end])
	return lowerBound
}

// Space Complexity is O(1)
func BinaryInsertionSort(input []int) []int {
	return MergeSort(input)
}

func binarySearch(input []int, value int) {

}
