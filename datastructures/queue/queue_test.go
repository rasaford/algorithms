package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBounded(t *testing.T) {
	tests := []struct {
		name string
		want Queue
		size int
	}{
		{
			"empty queue",
			&boundedQueue{
				array: make([]interface{}, 15+1),
				tail:  1,
				head:  1,
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBounded(tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnbounded(t *testing.T) {
	tests := []struct {
		name string
		want Queue
	}{
		{
			"empty queue",
			&unboundedQueue{
				array: make([]interface{}, 0, 8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnbounded(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boundedQueue_Enqueue(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Queue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewBounded(len(array)),
			args{array},
			args{reverse(array)},
			false,
		},
		{
			"queue overflow",
			NewBounded(len(array) - 1),
			args{array},
			args{reverse(array)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("sentinel")
			for _, v := range tt.args.value {
				if err = tt.q.Enqueue(v); ((err != nil) != tt.wantErr) && !tt.wantErr {
					t.Errorf("boundedQueue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, _ := tt.q.Dequeue(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("boundedQueue.Dequeue() got %v, want %v", got, v)
				}
			}
		})
	}
}

func Test_boundedQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		q       Queue
		want    interface{}
		wantErr bool
	}{
		{
			"empty queue",
			NewBounded(555),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("boundedQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) != tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boundedQueue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unboundedQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		q       Queue
		want    interface{}
		wantErr bool
	}{
		{
			"empty queue",
			NewUnbounded(),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("unboundedQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unboundedQueue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unboundedQueue_Enqueue(t *testing.T) {
	array := []interface{}{1, 2, 4, 5, 5, 6, 6, 4, 34, 35345, 45645, 45642, 11, 1, 3434, 33000}
	type args struct {
		value []interface{}
	}
	tests := []struct {
		name       string
		q          Queue
		args, want args
		wantErr    bool
	}{
		{
			"int queue",
			NewUnbounded(),
			args{array},
			args{reverse(array)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.value {
				if err := tt.q.Enqueue(v); (err != nil) != tt.wantErr {
					t.Errorf("unboundedQueue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			for _, v := range tt.want.value {
				if got, err := tt.q.Dequeue(); (err != nil) != tt.wantErr &&
					!reflect.DeepEqual(got, v) {
					t.Errorf("unboundedQueue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func reverse(input []interface{}) []interface{} {
	for i := 0; i < len(input)/2; i++ {
		opposite := len(input) - i - 1
		temp := input[opposite]
		input[opposite] = input[i]
		input[i] = temp
	}
	return input
}
