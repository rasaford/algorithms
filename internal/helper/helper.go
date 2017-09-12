package helper

import "math/rand"
import "fmt"
import "math"

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

// RandBetween generates random numbers in the range [a .. b] inclusive.
func RandBetween(a, b int) (int, error) {
	if a > b {
		return -1, fmt.Errorf("%d is larger than %d", a, b)
	}
	if a < 0 || b < 0 {
		return -1, fmt.Errorf("%f is less than 0", math.Min(float64(a), float64(b)))
	}
	return rand.Intn(b-a+1) + a, nil
}
