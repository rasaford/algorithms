package dynamic

import (
	"reflect"
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	type args struct {
		lengths []int
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 map[string]int
	}{
		{
			"book example",
			args{[]int{30, 35, 15, 5, 10, 20, 25}},
			map[string]int{str(1, 3): 7875, str(1, 5): 11875, str(2, 6): 10500, str(1, 4): 9375, str(2, 5): 7125, str(2, 3): 2625, str(3, 4): 750, str(4, 5): 1000, str(2, 4): 4375, str(4, 6): 3500, str(1, 2): 15750, str(5, 6): 5000, str(3, 5): 2500, str(3, 6): 5375, str(1, 6): 15125},
			map[string]int{str(2, 5): 3, str(2, 3): 2, str(3, 4): 3, str(5, 6): 5, str(2, 4): 3, str(3, 6): 3, str(4, 6): 5, str(1, 5): 3, str(2, 6): 3, str(1, 6): 3, str(1, 2): 1, str(1, 3): 1, str(3, 5): 3, str(4, 5): 4, str(1, 4): 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrices, store := MatrixChainOrder(tt.args.lengths)
			if !reflect.DeepEqual(matrices, tt.want) {
				t.Errorf("MatrixChainOrder() matrices = %v, want %v", matrices, tt.want)
			}
			if !reflect.DeepEqual(store, tt.want1) {
				t.Errorf("MatrixChainOrder() store = %v, want %v", store, tt.want1)
			}
		})
	}
}

func TestPrintOptimal(t *testing.T) {
	type args struct {
		matrixRow []int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"book example",
			args{[]int{30, 35, 15, 5, 10, 20, 25}},
			"15125",
			"((A_1(A_2A_3))((A_4A_5)A_6))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cost, parens := PrintOptimal(tt.args.matrixRow)
			if cost != tt.want {
				t.Errorf("PrintOptimal() cost = %v, want %v", cost, tt.want)
			}
			if parens != tt.want1 {
				t.Errorf("PrintOptimal() parens = %v, want %v", parens, tt.want1)
			}
		})
	}
}

func TestMatrixChainMultiply(t *testing.T) {
	type args struct {
		matrices [][][]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			"5x7 example",
			args{[][][]int{
				[][]int{
					[]int{1, 2, 3, 4, 5},
					[]int{6, 7, 8, 9, 10},
					[]int{11, 12, 13, 14, 15},
					[]int{16, 17, 18, 19, 20},
				},
				[][]int{
					[]int{1, 2},
					[]int{3, 4},
					[]int{5, 6},
					[]int{7, 8},
					[]int{9, 10},
				},
				[][]int{
					[]int{1, 2},
					[]int{3, 4},
				},
				[][]int{
					[]int{1, 2, 3, 4, 5, 6, 7},
					[]int{8, 9, 10, 11, 12, 13, 14},
				},
			}},
			[][]int{
				[]int{87, 104, 121, 138, 155, 172, 189},
				[]int{191, 228, 265, 302, 339, 376, 413},
				[]int{295, 352, 409, 466, 523, 580, 637},
				[]int{399, 476, 553, 630, 707, 784, 861},
				[]int{503, 600, 697, 794, 891, 988, 1085},
			},
			false,
		},
		{
			"invalid matrices",
			args{[][][]int{
				[][]int{
					[]int{1, 2, 3, 4, 5},
					[]int{6, 7, 8, 9, 10},
					[]int{11, 12, 13, 14, 15},
					[]int{16, 17, 18, 19, 20},
				},
				[][]int{
					[]int{1, 2},
					[]int{3, 4},
				},
			}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatrixChainMultiply(tt.args.matrices)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatrixChainMultiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixChainMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
