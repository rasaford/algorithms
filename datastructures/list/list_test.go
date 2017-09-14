package list

import (
	"reflect"
	"testing"
)

func TestNode_Next(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Prev(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *List
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Start(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Start(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_End(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.End(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("List.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		l       *List
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Remove(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("List.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name    string
		l       *List
		args    args
		want    *Node
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertAfter(t *testing.T) {
	type args struct {
		val  int
		node *Node
	}
	tests := []struct {
		name string
		l    *List
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.InsertAfter(tt.args.val, tt.args.node)
		})
	}
}
