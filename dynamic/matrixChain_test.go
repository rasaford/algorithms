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
			map[string]int{
				str(1, 3): 7875,
				str(1, 5): 11875,
				str(2, 6): 10500,
				str(1, 4): 9375,
				str(2, 5): 7125,
				str(2, 3): 2625,
				str(3, 4): 750,
				str(4, 5): 1000,
				str(2, 4): 4375,
				str(4, 6): 3500,
				str(1, 2): 15750,
				str(5, 6): 5000,
				str(3, 5): 2500,
				str(3, 6): 5375,
				str(1, 6): 15125,
			},
			map[string]int{
				str(2, 5): 3,
				str(2, 3): 2,
				str(3, 4): 3,
				str(5, 6): 5,
				str(2, 4): 3,
				str(3, 6): 3,
				str(4, 6): 5,
				str(1, 5): 3,
				str(2, 6): 3,
				str(1, 6): 3,
				str(1, 2): 1,
				str(1, 3): 1,
				str(3, 5): 3,
				str(4, 5): 4,
				str(1, 4): 3,
			},
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
