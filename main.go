package main

import (
	"fmt"

	"github.com/rasaford/algorithms/sorting"
)

func main() {
	res := sorting.MergeSort([]int{5, 4, 3, 2})
	fmt.Println(res)
}
