package sorting

import (
	"math/rand"
	"reflect"
	"testing"
)

type args struct {
	array []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		"reverse sorted",
		args{[]int{5, 4, 3, 2, 1}},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"random",
		args{[]int{-1503, 527, 1479, -837, -639, -1658, 825,
			-1364, 1864, -92, 1222, 687, 1954, 2033, -1704,
			-742, -373, 1173, 1829, 994, -241, 218, -1432,
			1682, -897, 1323, -2001, 1272, 1078, -9, 821,
			1912, 475, 783, -91, 332, -2007, -1033, -1283,
			914, -371, -1390, -1078, -1213, 497, 915, -1058,
			740, -1665, -935, -1259, -439, -1547, -873, 1192,
			1553, 712, -1286, -153, -85, -1277, 30, 189, 412}},
		[]int{-2007, -2001, -1704, -1665, -1658, -1547, -1503,
			-1432, -1390, -1364, -1286, -1283, -1277, -1259,
			-1213, -1078, -1058, -1033, -935, -897, -873,
			-837, -742, -639, -439, -373, -371, -241, -153,
			-92, -91, -85, -9, 30, 189, 218, 332, 412, 475,
			497, 527, 687, 712, 740, 783, 821, 825, 914, 915,
			994, 1078, 1173, 1192, 1222, 1272, 1323, 1479,
			1553, 1682, 1829, 1864, 1912, 1954, 2033},
	},
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func random(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(n*n) - n*n/2
	}
	return s
}
