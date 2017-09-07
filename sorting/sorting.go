package sorting

import (
	"math"
)

const inf = math.MaxInt32
const negInf = math.MaxInt32

// MergeSort is a more complex implmentation of sorting, it
// recursively splits the input and sorts the subarrays. Then
// it combines the inputs to form the original array. It therefore uses the
// divide and conquer paradigm.
//
// It runs in O(n lg n) time, where n := len(input)
func MergeSort(input []int) []int {
	array := clone(input)
	mergeSortRec(array, 0, len(array)-1)
	return array
}

func mergeSortRec(input []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSortRec(input, 0, mid)
		mergeSortRec(input, mid+1, high)
		merge(input, low, mid, high)
	}
}

func merge(input []int, low, mid, high int) {
	lower := mid - low + 1
	upper := high - mid
	left := make([]int, lower+1)
	right := make([]int, upper+1)
	for i := 0; i < lower; i++ {
		left[i] = input[low+i]
	}
	for i := 1; i <= upper; i++ {
		right[i-1] = input[mid+i]
	}
	left[lower] = inf
	right[upper] = inf
	i, j := 0, 0
	for k := low; k <= high; k++ {
		if left[i] <= right[j] {
			input[k] = left[i]
			i++
		} else {
			input[k] = right[j]
			j++
		}
	}
}

// BubbleSort is the naive implementation of sorting an array.
// It is easy to implement but not very efficinet.
//
// It runs in O(n^2) time, where n := len(input)
func BubbleSort(input []int) []int {
	array := clone(input)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				swap(&array[i], &array[j])
			}
		}
	}
	return array
}

// SelectionSort is a simple sorting algorithm that scans
// through the array and maintains a sorted version of it
// behind the read head.
//
// It runs in O(n^2) time, where n := len(input)
func SelectionSort(input []int) []int {
	array := clone(input)
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

func QuickSort(input []int) []int {
	return MergeSort(input)
}

func HeapSort(input []int) []int {
	return MergeSort(input)
}

func BinaryInsertionSort(input []int) []int {
	return MergeSort(input)
}

// swap uses pointers to two ints to swap them
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// clones the array and returns the copy
func clone(input []int) []int {
	array := make([]int, len(input))
	copy(array, input)
	return array
}
