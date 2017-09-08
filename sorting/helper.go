package sorting

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
