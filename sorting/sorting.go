package sorting

// MergeSort is a more complex implmentation of sorting, it
// recursively splits the input and sorts the subarrays. Then
// it combines the inputs to form the original array. It therefore uses the
// divide and conquer paradigm.
// It runs in O(n lg n) time, where n := len(input)
func MergeSort(input []int) []int {
	// TODO: Implement MergeSort
	return BubbleSort(input)
}

// BubbleSort is the naive implementation of sorting an array.
// It is easy to implement but not very efficinet.
// It runs in O(n^2) time, where n := len(input)
func BubbleSort(input []int) []int {
	array := make([]int, len(input))
	copy(array, input)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				swap(&array[i], &array[j])
			}
		}
	}
	return array
}

func SelectionSort(input []int) []int {
	// TODO: Implemente Selection Sort
	return BubbleSort(input)
}

// swap uses pointers to two ints to swap them
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
