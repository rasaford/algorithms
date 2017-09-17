package main

import (
	"github.com/rasaford/algorithms/datastructures/queue"
)

func main() {
	q := queue.NewDequeue(55)
	q.DequeueTail()
}
