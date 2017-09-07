package matrix

import (
	"fmt"
	"math"
)

func validateMatriciesStrassen(a, b [][]int) error {
	if err := validateMatricies(a, b); err != nil {
		return err
	}
	log := math.Log2(float64(len(a)))
	if log != float64(int(log)) {
		return fmt.Errorf("the size of the matricies is not a power of 2")
	}
	return nil
}

func validateMatricies(a, b [][]int) error {
	if a == nil || b == nil {
		return fmt.Errorf("one of the maticies is nil")
	}
	if len(a) != len(b) {
		return fmt.Errorf("the matricies are not of equal size")
	}
	for i, v := range a {
		if len(v) != len(b[i]) {
			return fmt.Errorf("the matricies are not of equal size")
		}
	}
	return nil
}

func newSquareMatrix(n int) [][]int {
	new := make([][]int, n)
	for i := range new {
		new[i] = make([]int, n)
	}
	return new
}
