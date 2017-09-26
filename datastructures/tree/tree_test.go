package tree

import (
	"reflect"
	"testing"
)

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
		t    *binaryTree
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
			if got := tt.t.Successor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
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
		t    *binaryTree
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
			if got := tt.t.Predecessor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
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
		// {
		// "random delete",
		// NewBinaryTree(),
		// []args{
		// 	{224756258, 945986064},
		// 	{353390024, -133757764},
		// 	{-161839742, 401199495},
		// 	{-932787398, -737619283},
		// 	{-865501367, -427538523},
		// 	{32668871, 673536688},
		// 	{-613613661, -256286734},
		// 	{-390717095, -66808549},
		// 	{-465930612, -444310378},
		// 	{384581414, -604402717},
		// 	{-637401328, -298776357},
		// 	{151769705, 778444435},
		// 	{-444283776, -435761876},
		// 	{542396464, -630108935},
		// 	{784550967, 422451192},
		// {51153718, -1012961415},
		// {-733734436, 230325008},
		// {1020573606, -903116467},
		// {203599914, -946781192},
		// {412369662, -426226797},
		// {-701655410, 88261267},
		// {94823371, -475651260},
		// {-165029390, 65682324},
		// {-529267745, -467977499},
		// {619774336, -296770470},
		// {817210134, -435698102},
		// {846885366, -864459624},
		// {1024171177, -914203118},
		// {-596378935, 388862715},
		// {-555092120, -404752468},
		// {929530628, 519366687},
		// {646510782, 494418333},
		// {-680913557, -153851995},
		// {852533738, 392245379},
		// {1028492960, 906693923},
		// {-878670260, -14727447},
		// {916947179, 976987895},
		// {-326516376, 409823271},
		// {452919754, 136965640},
		// {321026173, 111164597},
		// {549376800, -206580870},
		// {-793170686, 1043601310},
		// {851137418, -382071764},
		// {474911210, 310396820},
		// {-890087932, 364160179},
		// {263557055, -279832487},
		// {-565169277, 75767284},
		// {-671633883, -560835320},
		// {275088729, -801541980},
		// {-469589618, -192580225},
		// {-139774398, 268639528},
		// {107689692, 265447934},
		// {492161864, 709816187},
		// {-1072638413, 506953462},
		// {-214783233, -4578191},
		// {223291275, -194093270},
		// {-1010023232, -1069653241},
		// {-1067636439, 892969474},
		// {192917444, 127544314},
		// {677327448, 811774071},
		// {-89244097, 215103979},
		// {-1017337842, 742670256},
		// {-537529757, 304479447},
		// {-542311330, -700818737},
		// {198907996, 675157158},
		// {416264229, -1008624648},
		// {84203062, 1021503887},
		// {538509564, -442368074},
		// {543659705, -749549005},
		// {-309737439, 712816080},
		// {-575890599, 274522724},
		// {-3448211, -880820291},
		// {-1019638207, -231463984},
		// {191948717, 922583963},
		// {154805228, 190216253},
		// {-189488183, 112915528},
		// {-18022979, 983448541},
		// {638250482, -843142640},
		// {607812978, -229241733},
		// {-793680221, -665649565},
		// {515021944, 330801408},
		// {-862464245, 43766331},
		// {-859574001, -747660600},
		// {-910124481, -396837615},
		// {-730894080, -777809855},
		// {-380940657, 83911887},
		// {152152712, 27448617},
		// {395513081, 328651339},
		// {52612834, 331292592},
		// {464647548, 293441216},
		// {-1046198393, -1007852310},
		// {-863222123, -281080464},
		// {701054897, -327101040},
		// {-334330953, -530428839},
		// {-608873576, 118116187},
		// {-210301258, 13952339},
		// {-711504994, -362133903},
		// {704220225, 430114933},
		// {-949346128, 1071936843},
		// {-189965623, -833922367},
		// {602914806, -875920731},
		// {-958862992, 461055754},
		// {-535233937, 748683497},
		// {1017653576, -617270661},
		// {-1027498376, 956048469},
		// {-874089935, 313174788},
		// {-403972720, -110671935},
		// {-27403505, -896618067},
		// {369000196, -214344016},
		// {859584334, 966116829},
		// {-387983148, -1319642},
		// {-213819970, -1031203028},
		// {311468594, -153140430},
		// {-344463353, 832037173},
		// {-566232437, 569100809},
		// {-996959303, 488718437},
		// {270232099, 28105206},
		// {-918160162, 481528237},
		// {815710597, 1025989256},
		// {746251131, 713389641},
		// {-541499763, 887767728},
		// {-912600642, 719629934},
		// {277737693, 540608777},
		// {283475216, -865577186},
		// {-1041900290, 180033738},
		// {-926089019, 1070034862},
		// 	},
		// },
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
					t.Errorf("binaryTree.Delete(), don't want %v; got: %v", prev, res)
				}
			}
		})
	}
}
