package helper

// Swap uses pointers to two ints to swap them
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Clone clones the array and returns the copy
func Clone(input []int) []int {
	array := make([]int, len(input))
	copy(array, input)
	return array
}
