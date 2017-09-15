package stack

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{
			"empty stack",
			&Stack{
				array: make([]interface{}, 0, 8),
				top:   -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want bool
	}{
		{
			"empty stack",
			&Stack{
				array: make([]interface{}, 0, 8),
				top:   -1,
			},
			true,
		},
		{
			"nonempty stack",
			&Stack{
				array: make([]interface{}, 1, 8),
				top:   0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Stack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	a := 555
	g := struct {
		a int
		b string
	}{
		a: 5,
		b: "this is a test",
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
	}{
		{
			"int push",
			New(),
			args{a},
		},
		{
			"struct push",
			New(),
			args{g},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				tt.s.Push(tt.args.val)
				if got, _ := tt.s.Pop(); !reflect.DeepEqual(got, tt.args.val) {
					t.Errorf("Stack.Push() does not write the value")
				}
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    *interface{}
		wantErr bool
	}{
		{
			"pop form empty stack",
			New(),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) != tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
