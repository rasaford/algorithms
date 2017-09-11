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

func TestHeap_Insert(t *testing.T) {
	type args struct {
		element int
	}
	tests := []struct {
		name    string
		h       *Heap
		args    args
		wantErr bool
	}{
		{
			"insert into empty max heap",
			NewMaxHeap(make([]int, 0)),
			args{123123},
			false,
		},
		{
			"insert into non-empty max heap",
			NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			args{123123},
			false,
		},
		{
			"insert into non-empty min heap",
			NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			args{-123123},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Insert(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Queue.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := validateHeap(tt.h.heap, tt.h.compare); err != nil {
				t.Errorf("Queue.Insert() results in invalid heap")
			}
		})
	}
}

func TestHeap_Head(t *testing.T) {
	tests := []struct {
		name string
		h    *Heap
		want int
	}{
		{
			"get max element",
			NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			23,
		},
		{
			"get min element",
			NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Head(); got != tt.want {
				t.Errorf("Queue.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHead_ExtractHead(t *testing.T) {
	tests := []struct {
		name    string
		h       *Heap
		want    int
		wantErr bool
	}{
		{
			"max extract",
			NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			23,
			false,
		},
		{
			"min extract",
			NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			0,
			false,
		},
		{
			"heap too small",
			NewMinHeap([]int{}),
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.ExtractHead()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.ExtractHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.ExtractHead() = %v, want %v", got, tt.want)
			}
			if err := validateHeap(tt.h.heap, tt.h.compare); err != nil {
				t.Errorf("Queue.ExtractHead() results in invalid heap")
			}
		})
	}
}

func TestQueue_UpdateKey(t *testing.T) {
	type args struct {
		index int
		key   int
	}
	tests := []struct {
		name    string
		h       *Heap
		args    args
		wantErr bool
	}{
		{
			"root increase",
			NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			args{
				9, 555,
			},
			false,
		},
		{
			"new key smaller",
			NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			args{
				9, -55555,
			},
			true,
		},
		{
			"new key larger",
			NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			args{
				9, 55555,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UpdateKey(tt.args.index, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Queue.IncreaseKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := validateHeap(tt.h.heap, tt.h.compare); err != nil {
				t.Errorf("Queue.IncreaseKey(%d, %d) results in invalid heap %v", tt.args.index, tt.args.key, tt.h.heap)
			}
		})
	}
}
