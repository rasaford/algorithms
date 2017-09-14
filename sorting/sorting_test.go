package sorting

import (
	"math"
	"reflect"
	"testing"

	"github.com/rasaford/algorithms/internal/helper"
)

type args struct {
	input []int
}

var rand = helper.GenerateRandom(1<<12, math.MinInt32, math.MaxInt32)
var tests = []struct {
	name string
	args args
	want []int
}{
	{
		"reverse sorted",
		args{[]int{5, 4, 3, 2, 1, -55}},
		[]int{-55, 1, 2, 3, 4, 5},
	},
	{
		"short random",
		args{[]int{44, 534, -1, -223234, 3423,
			-5, -55555, -234, -111, 23423409}},
		[]int{-223234, -55555, -234, -111, -5,
			-1, 44, 534, 3423, 23423409},
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
	{
		"generated random",
		args{rand},
		QuickSort(rand),
	},
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryInsertionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryInsertionSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryInsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountingSort(t *testing.T) {
	randPos := helper.GenerateRandom(1<<12, 0, 1<<20)
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"reverse sorted",
			args{[]int{5, 4, 3, 2, 1, 55}},
			[]int{1, 2, 3, 4, 5, 55},
			false,
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
			true,
		},
		{
			"generated random",
			args{randPos},
			QuickSort(randPos),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := CountingSort(tt.args.input); !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v error: %v", got, tt.want, err)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type args struct {
		input       []int
		ithSmallest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"reverse sorted",
			args{[]int{5, 4, 3, 2, 1, 55}, 4},
			4,
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
				1553, 712, -1286, -153, -85, -1277, 30, 189, 412},
				15},
			-85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.input, tt.args.ithSmallest); got != tt.want {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
