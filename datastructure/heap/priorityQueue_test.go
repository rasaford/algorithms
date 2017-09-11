package heap

import (
	"testing"
)

func TestQueue_Insert(t *testing.T) {
	type args struct {
		element int
	}
	tests := []struct {
		name    string
		q       *queue
		args    args
		wantErr bool
	}{
		{
			"insert into empty max heap",
			&queue{
				NewMaxHeap(make([]int, 0)),
			},
			args{123123},
			false,
		},
		{
			"insert into non-empty max heap",
			&queue{
				NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			args{123123},
			false,
		},
		{
			"insert into non-empty min heap",
			&queue{
				NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			args{-123123},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Insert(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Queue.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := validateHeap(tt.q.heap, tt.q.compare); err != nil {
				t.Errorf("Queue.Insert() results in invalid heap")
			}
		})
	}
}

func TestQueue_Head(t *testing.T) {
	tests := []struct {
		name string
		q    *queue
		want int
	}{
		{
			"get max element",
			&queue{
				NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			23,
		},
		{
			"get min element",
			&queue{
				NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Head(); got != tt.want {
				t.Errorf("Queue.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_ExtractHead(t *testing.T) {
	tests := []struct {
		name    string
		q       *queue
		want    int
		wantErr bool
	}{
		{
			"max extract",
			&queue{
				NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			23,
			false,
		},
		{
			"min extract",
			&queue{
				NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			0,
			false,
		},
		{
			"heap too small",
			&queue{
				NewMinHeap([]int{}),
			},
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ExtractHead()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.ExtractHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.ExtractHead() = %v, want %v", got, tt.want)
			}
			if err := validateHeap(tt.q.heap, tt.q.compare); err != nil {
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
		q       *queue
		args    args
		wantErr bool
	}{
		{
			"root increase",
			&queue{
				NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			args{
				9, 555,
			},
			false,
		},
		{
			"new key smaller",
			&queue{
				NewMaxHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			args{
				9, -55555,
			},
			true,
		},
		{
			"new key larger",
			&queue{
				NewMinHeap([]int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23}),
			},
			args{
				9, 55555,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.UpdateKey(tt.args.index, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Queue.IncreaseKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := validateHeap(tt.q.heap, tt.q.compare); err != nil {
				t.Errorf("Queue.IncreaseKey(%d, %d) results in invalid heap", tt.args.index, tt.args.key)
			}
		})
	}
}
