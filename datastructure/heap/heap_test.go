package heap

import (
	"fmt"
	"testing"
)

type args struct {
	array []int
}

var tests = []struct {
	name string
	args args
}{
	{
		"reverse sorted",
		args{[]int{5, 4, 3, 2, 1, 0}},
	},
	{
		"book example",
		args{[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}},
	},
}

func TestNewMaxHeap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewMaxHeap(tt.args.array)
			comparator := func(a, b int) bool {
				if a > b {
					return true
				}
				return false
			}
			if err := validateHeap(tt.args.array, comparator); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestNewMinHeap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewMinHeap(tt.args.array)
			comparator := func(a, b int) bool {
				if a < b {
					return true
				}
				return false
			}
			if err := validateHeap(tt.args.array, comparator); err != nil {
				t.Error(err)
			}
		})
	}
}

func validateHeap(array []int, compare func(int, int) bool) error {
	for i := range array {
		left := left(i)
		right := right(i)
		if left <= len(array)-1 && compare(array[left], array[i]) ||
			right <= len(array)-1 && compare(array[right], array[i]) {
			return fmt.Errorf("The Heap condidtion is not valid on node %d", i)
		}
	}
	return nil
}
