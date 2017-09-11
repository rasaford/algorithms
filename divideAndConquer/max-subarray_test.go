package divideAndConquer

import (
	"reflect"
	"testing"
)

var cases = []struct {
	in, want []int
	sum      int
}{
	{
		[]int{13, -3, -25, 20,
			-3, -16, -23, 18,
			20, -7, 12, -5,
			-22, 15, -4, 7},
		[]int{18, 20, -7, 12},
		43,
	},
	{
		[]int{-5, -9, 20, -1,
			20, -5, -2000},
		[]int{20, -1, 20},
		39,
	},
	{
		[]int{-5, -40, 20, -10,
			5, -1, -1, -1,
			5, 300},
		[]int{20, -10, 5, -1,
			-1, -1, 5, 300},
		317,
	},
}

func TestMaxSubArrayRec(t *testing.T) {
	for _, c := range cases {
		res, sum := MaxSubArrayRec(c.in)
		if !reflect.DeepEqual(res, c.want) {
			t.Errorf("Arrays not equal: Want %v but got %v", c.want, res)
		}
		if c.sum != sum {
			t.Errorf("Sum not equal: Want %v but got %v", c.sum, sum)
		}
	}
}

func TestMaxSubArrayLin(t *testing.T) {
	for _, c := range cases {
		res, sum := MaxSubArrayLin(c.in)
		if !reflect.DeepEqual(res, c.want) {
			t.Errorf("Arrays not equal: Want %v but got %v", c.want, res)
		}
		if c.sum != sum {
			t.Errorf("Sum not equal: Want %v but got %v", c.sum, sum)
		}
	}
}
