package list

import (
	"math"
	"reflect"
	"testing"
)

func TestList_Start(t *testing.T) {
	list := New()
	list2 := New()
	for i := 0; i < 10; i++ {
		list2.Insert(i)
	}
	res1, _ := list.Search(math.MinInt32)
	res2, _ := list2.Search(0)
	tests := []struct {
		name string
		l    *List
		want *Node
	}{
		{
			"empty list",
			list,
			res1,
		},
		{
			"nonempty list",
			list2,
			res2,
		},
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
	list := New()
	list2 := New()
	for i := 0; i < 10; i++ {
		list2.Insert(i)
	}
	res1, _ := list.Search(math.MinInt32)
	res2, _ := list2.Search(9)
	tests := []struct {
		name string
		l    *List
		want *Node
	}{
		{
			"empty list",
			list,
			res1,
		},
		{
			"nonempty list",
			list2,
			res2,
		},
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
	list := New()
	list2 := New()
	for i := 0; i < 10; i++ {
		list2.Insert(i)
	}
	tests := []struct {
		name string
		l    *List
		want int
	}{
		{
			"empty list",
			list,
			0,
		},
		{
			"nonempty list",
			list2,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Search(t *testing.T) {
	list := New()
	for i := 0; i < 100; i++ {
		list.Insert(i)
	}
	res5, _ := list.Search(5)
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
		{
			"search empty list",
			New(),
			args{0},
			nil,
			true,
		},
		{
			"search contained",
			list,
			args{5},
			res5,
			false,
		},
		{
			"search not contained",
			list,
			args{-5},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Search(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Delete(t *testing.T) {
	list := New()
	for i := 0; i < 100; i++ {
		list.Insert(i)
	}
	del1, _ := list.Search(55)
	type args struct {
		node *Node
	}
	tests := []struct {
		name    string
		l       *List
		args    args
		wantErr bool
	}{
		{
			"delete from empty list",
			New(),
			args{nil},
			true,
		},
		{
			"delete contained item",
			list,
			args{del1},
			false,
		},
		{
			"delete node not contained in list",
			list,
			args{&Node{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Delete(tt.args.node); (err != nil) != tt.wantErr {
				t.Fatalf("List.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.args.node != nil && !tt.wantErr {
				_, err := tt.l.Search(tt.args.node.value)
				if err == nil {
					t.Errorf("List.Delete() did not delete the node: %v, %v", tt.args.node, err)
				}
			}
		})
	}
}

func TestConcat(t *testing.T) {
	listA := New()
	listB := New()
	for i := 0; i < 1000; i++ {
		listA.Insert(i)
		listB.Insert(i * i)
	}
	type args struct {
		a *List
		b *List
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		{
			"empty lists",
			args{New(), New()},
			New(),
		},
		{
			"first empty",
			args{New(), listB},
			listB,
		},
		{
			"second empty",
			args{listA, New()},
			listA,
		},
		{
			"nonempty lists",
			args{listA, listB},
			listA,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
