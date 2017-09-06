package main

import (
	"fmt"

	divide "github.com/rasaford/algorithms/divideAndConquer"
)

func main() {
	divide.Divide()
	input := []int{13, -3, -25, 20,
		-3, -16, -23, 18,
		20, -7, 12, -5,
		-22, 15, -4, 7}
	res, _ := divide.MaxSubArrayLin(input)
	fmt.Println(res)
}
