package sorting

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
	array, work := clone(input), clone(input)
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
