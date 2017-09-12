package main

import (
	"fmt"

	"github.com/rasaford/algorithms/internal/helper"
)

func main() {
	min, max := 1<<31, -1<<31
	for i := 0; i < 1<<15; i++ {
		rand, _ := helper.RandBetween(4, 10)
		if rand > max {
			max = rand
		}
		if rand < min {
			min = rand
		}
	}
	fmt.Printf("min: %v, max %v\n", min, max)
}
