package tree

import (
	"reflect"
	"testing"
)

//		  1
// 	   /    \
//    2      5
//  /  \   /  \
// 3   4  6    7
// var bTree = &binaryTree{
// 	root: &binaryNode{
// 		1,
// 		&binaryNode{
// 			2,
// 			&binaryNode{3, nil, nil},
// 			&binaryNode{4, nil, nil},
// 		},
// 		&binaryNode{
// 			5,
// 			&binaryNode{6, nil, nil},
// 			&binaryNode{7, nil, nil},
// 		},
// 	},
// }

// func TestBNode_walkRecursive(t *testing.T) {
// 	str := ""
// 	type args struct {
// 		do func(int)
// 	}
// 	tests := []struct {
// 		name string
// 		t    *binaryTree
// 		args args
// 		want string
// 	}{
// 		{
// 			"walk the tree",
// 			bTree,
// 			args{func(a int) {
// 				str = fmt.Sprintf("%s %d", str, a)
// 			}},
// 			" 1 2 3 4 5 6 7",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.t.root.walkRecursive(tt.args.do)
// 			if str != tt.want {
// 				t.Errorf("BNode.walkRecursive(), want %s; got: %s", tt.want, str)
// 			}
// 		})
// 	}
// }

// func TestBNode_walkIterative(t *testing.T) {
// 	str := ""
// 	type args struct {
// 		do func(int)
// 	}
// 	tests := []struct {
// 		name string
// 		t    *binaryTree
// 		args args
// 		want string
// 	}{
// 		{
// 			"walk the tree",
// 			bTree,
// 			args{func(a int) {
// 				str = fmt.Sprintf("%s %d", str, a)
// 			}},
// 			" 1 2 3 4 5 6 7",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.t.root.walkIterative(tt.args.do)
// 			if str != tt.want {
// 				t.Errorf("BNode.walkIterative(), want %s; got: %s", tt.want, str)
// 			}
// 		})
// 	}
// }

// func Test_binaryTree_String(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		t    *binaryTree
// 		want string
// 	}{
// 		{
// 			"walk the tree",
// 			bTree,
// 			"1 2 3 4 5 6 7",
// 		},
// 		{
// 			"print empty tree",
// 			NewBinaryTree().(*binaryTree),
// 			"",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.t.String(); got != tt.want {
// 				t.Errorf("binaryTree.String() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_binaryTree_Insert(t *testing.T) {
	type args struct {
		key   int
		value interface{}
	}
	tests := []struct {
		name string
		t    Tree
		args []args
	}{
		{
			"1..10",
			NewBinaryTree(),
			[]args{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
				{6, 6},
				{7, 7},
				{8, 8},
				{9, 9},
				{10, 10},
			},
		},
		{
			"duplicated keys",
			NewBinaryTree(),
			[]args{
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.t.Insert(arg.key, arg.value)
			}
			for _, arg := range tt.args {
				if res := tt.t.Search(arg.key); res != arg.value {
					t.Errorf("BNode.Insert(), want %v; got: %v", arg.value, res)
				}
			}
		})
	}
}

func Test_binaryTree_Search(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(3, 3)
	tree.Insert(2, 2)
	tree.Insert(4, 4)
	tree.Insert(1, 1)
	tree.Insert(5, 5)
	type args struct {
		key int
	}
	tests := []struct {
		name string
		t    Tree
		args args
		want interface{}
	}{
		{
			"empty tree",
			NewBinaryTree(),
			args{55},
			nil,
		},
		{
			"node contained",
			tree,
			args{5},
			5,
		},
		{
			"node not contained",
			tree,
			args{-5},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTree_successor(t *testing.T) {
	tree := NewBinaryTree().(*binaryTree)
	tree.Insert(3, 3)
	tree.Insert(2, 2)
	tree.Insert(4, 4)
	tree.Insert(1, 1)
	tree.Insert(5, 5)
	type args struct {
		node *binaryNode
	}
	tests := []struct {
		name string
		t    Tree
		args args
		want *binaryNode
	}{
		{
			"1",
			tree,
			args{tree.search(1)},
			tree.search(2),
		},
		{
			"2",
			tree,
			args{tree.search(2)},
			tree.search(3),
		},
		{
			"3",
			tree,
			args{tree.search(3)},
			tree.search(4),
		},
		{
			"4",
			tree,
			args{tree.search(4)},
			tree.search(5),
		},
		{
			"not contained, large key",
			tree,
			args{tree.search(5)},
			nil,
		},
		{
			"not contained, nil key",
			tree,
			args{tree.search(555)},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.successor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTree.Successor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTree_predecessor(t *testing.T) {
	tree := NewBinaryTree().(*binaryTree)
	tree.Insert(3, 3)
	tree.Insert(2, 2)
	tree.Insert(4, 4)
	tree.Insert(1, 1)
	tree.Insert(5, 5)
	type args struct {
		node *binaryNode
	}
	tests := []struct {
		name string
		t    Tree
		args args
		want *binaryNode
	}{
		{
			"2",
			tree,
			args{tree.search(2)},
			tree.search(1),
		},
		{
			"3",
			tree,
			args{tree.search(3)},
			tree.search(2),
		},
		{
			"4",
			tree,
			args{tree.search(4)},
			tree.search(3),
		},
		{
			"5",
			tree,
			args{tree.search(5)},
			tree.search(4),
		},
		{
			"not contained, small key",
			tree,
			args{tree.search(1)},
			nil,
		},
		{
			"not contained, nil key",
			tree,
			args{tree.search(555)},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.predecessor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTree.Successor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTree_Delete(t *testing.T) {
	type args struct {
		key   int
		value interface{}
	}
	tests := []struct {
		name string
		t    Tree
		args []args
	}{
		{
			"root delete",
			NewBinaryTree(),
			[]args{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
				{6, 6},
				{7, 7},
				{8, 8},
				{9, 9},
				{10, 10},
			},
		},
		{
			"non root delete",
			NewBinaryTree(),
			[]args{
				{6, 6},
				{7, 7},
				{8, 8},
				{9, 9},
				{10, 10},
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
			},
		},
		{
			"duplicated keys",
			NewBinaryTree(),
			[]args{
				{1, 2},
				{1, 3},
				{1, 4},
				{1, 5},
				{1, 6},
				{1, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.t.Insert(arg.key, arg.value)
			}
			for _, arg := range tt.args {
				prev := tt.t.Search(arg.key)
				tt.t.Delete(arg.key)
				if res := tt.t.Search(arg.key); reflect.DeepEqual(prev, res) {
					t.Errorf("binaryTree.Delete(), want %v; got: %v", nil, res)
				}
			}
		})
	}
}
