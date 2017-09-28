package dynamic

import (
	"fmt"
	"math"
	"strconv"

	"github.com/rasaford/algorithms/divideAndConquer/matrix"
	"github.com/rasaford/algorithms/internal/helper"
)

func MatrixChainMultiply(matrices [][][]int) ([][]int, error) {
	rowCount, err := getRowCount(matrices)
	if err != nil {
		return nil, err
	}
	_, store := MatrixChainOrder(rowCount)
	return multiplyPairs(matrices, store, 1, len(matrices)-1), nil
}

func getRowCount(matrices [][][]int) ([]int, error) {
	rowCount := make([]int, len(matrices))
	for i := range rowCount {
		rowCount[i] = len(matrices[i])
		if i != 0 && len(matrices[i-1][0]) != len(matrices[i]) {
			return nil, fmt.Errorf("column of matrix %d is not equal to row of matrix %d", i-1, i)
		}
	}
	return rowCount, nil
}

func multiplyPairs(matrices [][][]int, store map[string]int, i, j int) [][]int {
	if i == j {
		return matrices[i]
	}
	split := store[str(i, j)]
	// no error can occur in MatrixMultiply because it has already been checked that
	// the column count of the first matrix is the same as thr row count on the second.
	res, _ := matrix.Multiply(
		multiplyPairs(matrices, store, i, split),
		multiplyPairs(matrices, store, split+1, j))
	return res
}

// MatrixChainOrder takes an array of row counts for n matrices and computes the best way to
// parenthesize multiplying all of them together.
// A chain A: A_1 * A_2 * A_3 can for example be parenthesized as:
// A:((A_1 * A_2) * A_3). If the matrices A_i are non-square the amount of scalar
// multiplications required can be greatly reduced.
// The product of a p x q and a q x r matrix has the dimensions: p x r
//
// It runs in O(n^3) time, wehre n := len(matrixRow)
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

func PrintOptimal(matrixRow []int) (string, string) {
	matrix, store := MatrixChainOrder(matrixRow)
	n := len(matrixRow) - 1
	cost := strconv.Itoa(matrix[str(1, n)])
	parens := printParens(store, 1, n)
	return cost, parens
}

func printParens(store map[string]int, i, j int) string {
	if i == j {
		return fmt.Sprintf("A_%d", i)
	}
	return helper.Concat("(",
		printParens(store, i, store[str(i, j)]),
		printParens(store, store[str(i, j)]+1, j),
		")")
}

func str(a, b int) string { return fmt.Sprintf("%d-%d", a, b) }
