package main

import (
	"fmt"

	"github.com/rasaford/algorithms/divideAndConquer/matrix"
)

func main() {
	a := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	b := [][]int{
		[]int{2, 0},
		[]int{1, 2},
	}
	res, _ := matrix.SquareMultiplyStrassen(a, b)
	fmt.Println(res)
}
