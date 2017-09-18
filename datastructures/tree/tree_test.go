package tree

import (
	"fmt"
	"testing"
)

//		  1
// 	   /    \
//    2      5
//  /  \   /  \
// 3   4  6    7
var bTree = &binaryTree{
	root: &binaryNode{
		1,
		&binaryNode{
			2,
			&binaryNode{3, nil, nil},
			&binaryNode{4, nil, nil},
		},
		&binaryNode{
			5,
			&binaryNode{6, nil, nil},
			&binaryNode{7, nil, nil},
		},
	},
}

func TestBNode_walkRecursive(t *testing.T) {
	str := ""
	type args struct {
		do func(int)
	}
	tests := []struct {
		name string
		t    *binaryTree
		args args
		want string
	}{
		{
			"walk the tree",
			bTree,
			args{func(a int) {
				str = fmt.Sprintf("%s %d", str, a)
			}},
			" 1 2 3 4 5 6 7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.root.walkRecursive(tt.args.do)
			if str != tt.want {
				t.Errorf("BNode.walkRecursive(), want %s; got: %s", tt.want, str)
			}
		})
	}
}

func TestBNode_walkIterative(t *testing.T) {
	str := ""
	type args struct {
		do func(int)
	}
	tests := []struct {
		name string
		t    *binaryTree
		args args
		want string
	}{
		{
			"walk the tree",
			bTree,
			args{func(a int) {
				str = fmt.Sprintf("%s %d", str, a)
			}},
			" 1 2 3 4 5 6 7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.root.walkIterative(tt.args.do)
			if str != tt.want {
				t.Errorf("BNode.walkIterative(), want %s; got: %s", tt.want, str)
			}
		})
	}
}

func Test_binaryTree_String(t *testing.T) {
	tests := []struct {
		name string
		t    *binaryTree
		want string
	}{
		{
			"walk the tree",
			bTree,
			"1 2 3 4 5 6 7",
		},
		{
			"print empty tree",
			NewBinaryTree().(*binaryTree),
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("binaryTree.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
