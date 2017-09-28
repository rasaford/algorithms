package dynamic

import (
	"fmt"
	"math"
)

// MatrixChainOrder takes an array of row counts for n matrices and computes the best way to
// parenthesize multiplying all of them together.
// A chain A: A_1 * A_2 * A_3 can for example be parenthesized as:
// A:((A_1 * A_2) * A_3). If the matrices A_i are non-square the amount of scalar
// multiplications required can be greatly reduced.
// The product of a p x q and a q x r matrix has the dimensions: p x r
//
// It runs in O(n^2) time, wehre n := len(matrixRow)
func MatrixChainOrder(matrixRow []int) (map[string]int, map[string]int) {
	n := len(matrixRow) - 1
	matrix, store := make(map[string]int), make(map[string]int)
	// if less than 2 matices need to be multiplied together the result is trivial.
	for l := 2; l <= n; l++ { // length of matrix chain
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			matrix[str(i, j)] = math.MaxInt32
			for split := i; split < j; split++ { // consider all ways of spliting the range i <= split < j
				q := matrix[str(i, split)] + matrix[str(split+1, j)] +
					matrixRow[i-1]*matrixRow[split]*matrixRow[j]
				if q < matrix[str(i, j)] { // set the result to the min of all possibilities
					matrix[str(i, j)] = q
					store[str(i, j)] = split
				}
			}
		}
	}
	return matrix, store
}

func str(a, b int) string { return fmt.Sprintf("%d-%d", a, b) }
