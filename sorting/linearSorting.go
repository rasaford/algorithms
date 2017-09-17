package sorting

import (
	"fmt"

	"github.com/rasaford/algorithms/internal/helper"
)

// CountingSort assumes that the given integers are in a small range (specifically O(n)).
// Based on this assumption the values frequency is inserted into a working array at the index corresponding to their value.
// Then each index on the working array is set to be the sum of the previous values.
// This is then used as a mapping from the input array to the output.
// All values of the input have to be in the range [0..k]
//
// It runs in O(n+k) time with n := len(input) k := max(input)
// Space complexity is O(max(n,k))
func CountingSort(input []int) ([]int, error) {
	min, max := helper.FindMinMax(input)
	if min < 0 {
		return nil, fmt.Errorf("all values in the slice have to be positive")
	}
	working := make([]int, max+1)
	return countingSort(input, working, max), nil
}

func countingSort(input, working []int, max int) []int {
	out := make([]int, len(input))
	for i := range input {
		working[input[i]]++
	}
	for i := 1; i < len(working); i++ {
		working[i] += working[i-1]
	}
	for i := len(input) - 1; i >= 0; i-- {
		out[working[input[i]]-1] = input[i]
		working[input[i]]--
	}
	return out
}

// TODO: Implement BucketSort
func BucketSort(input []int) []int {
	return QuickSort(input)
}
