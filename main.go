package main

import (
	"fmt"

	"github.com/rasaford/algorithms/sorting"
)

func main() {
	array := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	res := sorting.HeapSort(array)
	fmt.Println(res)
}
