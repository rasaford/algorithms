package divideAndConquer

import (
	"testing"
)

func TestMaxSubArrayRec(t *testing.T) {
	cases := []struct {
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
			[]int{-5, -9, 20, -1, 20, -5, -2000},
			[]int{20, -1, 20},
			39,
		},
	}
	for _, c := range cases {
		res, sum := MaxSubArrayRec(c.in)
		if !sliceEqual(res, c.want) {
			t.Errorf("Arrays not equal: Want %v but got %v", c.want, res)
		}
		if c.sum != sum {
			t.Errorf("Sum not equal: Want %v but got %v", c.sum, sum)
		}
	}
}

func TestMaxSubArrayLin(t *testing.T) {
	cases := []struct {
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
			[]int{-5, -9, 20, -1, 20, -5, -2000},
			[]int{20, -1, 20},
			39,
		},
	}
	for _, c := range cases {
		res, sum := MaxSubArrayLin(c.in)
		if !sliceEqual(res, c.want) {
			t.Errorf("Arrays not equal: Want %v but got %v", c.want, res)
		}
		if c.sum != sum {
			t.Errorf("Sum not equal: Want %v but got %v", c.sum, sum)
		}
	}
}

func sliceEqual(a, b []int) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
