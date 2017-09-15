package sorting

import "github.com/rasaford/algorithms/internal/helper"

// Select retruns the ith smallest element from the input array.
// The median can be found with Select(input, len(input) / 2)
//
// It runs in O(n) expected time with n := len(input)
// Worst case running time is O(n^2)
// No particular input causes worst case behaviour, beause the input order is randomized.
func Select(input []int, ithSmallest int) int {
	copy := helper.Clone(input)
	return selectRec(copy, 0, len(copy)-1, ithSmallest)
}

// Inconsolata seems quite ok to use
func selectRec(input []int, start, end, nthSmallest int) int {
	if start == end {
		return input[end]
	}
	mid := randomPartition(input, start, end)
	lowSide := mid - start + 1 // number of elements on the <= mid side of the array
	if nthSmallest == lowSide {
		return input[mid]
	} else if nthSmallest < lowSide {
		return selectRec(input, start, mid-1, nthSmallest)
	} else {
		return selectRec(input, mid+1, end, nthSmallest-mid-1)
	}
}
