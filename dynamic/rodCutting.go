package dynamic

import (
	"fmt"
	"math"

	"github.com/rasaford/algorithms/internal/helper"
)

// CutRod takes a list of prices for various lengths of rods and returns
// the maximum achievable price for a rod of length n if all cuts are free.
// The prices table has the length of the rod as an index into the array and
// the corresponding price as the value for that index.
func CutRod(prices []int, length int) int {
	if length == 0 {
		return 0
	}
	q := math.MinInt32 // -infinity
	for i := 1; i <= length; i++ {
		q = helper.Max(q, prices[i]+CutRod(prices, length-i))
	}
	return q
}

// CutRod takes a list of prices for various lengths of rods and returns
// the maximum achievable price for a rod of length n if all cuts are free.
// The prices table has the length of the rod as an index into the array and
// the corresponding price as the value for that index.
func CutRodMemoized(prices []int, length int) int {
	store := make([]int, length+1)
	for i := range store {
		store[i] = math.MinInt32
	}
	return CutRodMemAux(prices, store, length)
}

func CutRodMemAux(prices, store []int, length int) int {
	if store[length] >= 0 {
		return store[length]
	}
	q := 0
	if length != 0 {
		q = math.MinInt32
		for i := 1; i <= length; i++ {
			q = helper.Max(q, prices[i]+CutRodMemAux(prices, store, length-i))
		}
	}
	store[length] = q
	return q
}

func CutRodBottomUp(prices []int, length int) int {
	store := make([]int, length+1)
	for i := 1; i <= length; i++ {
		q := math.MinInt32
		for j := 1; j <= i; j++ {
			q = helper.Max(q, prices[j]+store[i-j])
		}
		store[i] = q
	}
	return store[length]
}

func CutRodBottomUpExtended(prices []int, length int) ([]int, []int) {
	store, cuts := make([]int, length+1), make([]int, length+1)
	for i := 1; i <= length; i++ {
		q := math.MinInt32
		for j := 1; j <= i; j++ {
			if r := prices[j] + store[i-j]; q < r {
				q = r
				cuts[i] = j
			}
		}
		store[i] = q
	}
	return store, cuts
}

func CutRodPrint(prices []int, length int) int {
	store, cuts := CutRodBottomUpExtended(prices, length)
	n := length
	for n > 0 {
		fmt.Printf("%d ", cuts[n])
		n -= cuts[n]
	}
	fmt.Println()
	return store[length]
}
