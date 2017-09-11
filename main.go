package main

import (
	"fmt"

	"github.com/rasaford/algorithms/datastructure/heap"
)

func main() {
	array := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	h := heap.NewMinPriorityQueue(array)
	for range array {
		fmt.Println(h.ExtractHead())
	}
}
