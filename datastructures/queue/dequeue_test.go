package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDequeue(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want Dequeue
	}{
		{
			"min",
			args{2},
			&deQueue{
				make([]interface{}, 2+1),
				1,
				1,
			},
		},
		{
			"nonempty",
			args{51},
			&deQueue{
				make([]interface{}, 51+1),
				25,
				25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDequeue(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deQueue_EnqueueHead(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Dequeue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewDequeue(len(array)),
			args{array},
			args{reverse(array)},
			false,
		},
		{
			"queue overflow",
			NewDequeue(len(array) - 1),
			args{array},
			args{reverse(array)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("sentinel")
			for _, v := range tt.args.value {
				if err = tt.q.EnqueueHead(v); ((err != nil) != tt.wantErr) && !tt.wantErr {
					t.Errorf("Dequeue.EnqueueHead() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, _ := tt.q.DequeueHead(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("Dequeue.DequeueHead() got %v, want %v", got, v)
				}
			}
		})
	}
}

func Test_deQueue_EnqueueTail(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Dequeue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewDequeue(len(array)),
			args{array},
			args{reverse(array)},
			false,
		},
		{
			"queue overflow",
			NewDequeue(len(array) - 1),
			args{array},
			args{reverse(array)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("sentinel")
			for _, v := range tt.args.value {
				if err = tt.q.EnqueueTail(v); ((err != nil) != tt.wantErr) && !tt.wantErr {
					t.Errorf("Dequeue.EnqueueTail() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, _ := tt.q.DequeueTail(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("Dequeue.DequeueTail() got %v, want %v", got, v)
				}
			}
		})
	}
}

func Test_deQueue_DequeueHead(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Dequeue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewDequeue(len(array)),
			args{array},
			args{array},
			false,
		},
		{
			"queue overflow",
			NewDequeue(len(array) - 1),
			args{array},
			args{array},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("sentinel")
			for _, v := range tt.args.value {
				if err = tt.q.EnqueueTail(v); ((err != nil) != tt.wantErr) && !tt.wantErr {
					t.Errorf("Dequeue.EnqueueTail() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, _ := tt.q.DequeueHead(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("Dequeue.DequeueHead() got %v, want %v", got, v)
				}
			}
		})
	}
}

func Test_deQueue_DequeueTail(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Dequeue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewDequeue(len(array)),
			args{array},
			args{array},
			false,
		},
		{
			"queue overflow",
			NewDequeue(len(array) - 1),
			args{array},
			args{array},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("sentinel")
			for _, v := range tt.args.value {
				if err = tt.q.EnqueueHead(v); ((err != nil) != tt.wantErr) && !tt.wantErr {
					t.Errorf("Dequeue.EnqueueHead() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, _ := tt.q.DequeueTail(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("Dequeue.DequeueTail() got %v, want %v", got, v)
				}
			}
		})
	}
}
