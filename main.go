package main

import (
	"fmt"

	divide "github.com/rasaford/algorithms/divideAndConquer"
)

func main() {
	input := []int{13, -3, -25, 20,
		-3, -16, -23, 18,
		20, -7, 12, -5,
		-22, 15, -4, 7}
	input = []int{-5, -9, 20, -1, 20, -5, -2000, -4}
	fmt.Println(divide.MaxSubArrayRec(input))
	fmt.Println(divide.MaxSubArrayLin(input))
}
