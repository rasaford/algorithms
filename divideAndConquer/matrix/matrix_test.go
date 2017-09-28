package matrix

import (
	"reflect"
	"testing"
)

type args struct {
	a [][]int
	b [][]int
}

func TestSquareMultiply(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			"2x2 matrix",
			args{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
				[][]int{
					[]int{2, 0},
					[]int{1, 2},
				},
			},
			[][]int{
				[]int{4, 4},
				[]int{10, 8},
			},
			false,
		},
		{
			"2x2 matrix2",
			args{[][]int{
				[]int{1, 3},
				[]int{7, 5},
			},
				[][]int{
					[]int{6, 8},
					[]int{4, 2},
				},
			},
			[][]int{
				[]int{18, 14},
				[]int{62, 66},
			},
			false,
		},
		{
			"3x3 matrix",
			args{[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
				[][]int{
					[]int{9, 7, 8},
					[]int{6, 5, 4},
					[]int{3, 2, 1},
				},
			},
			[][]int{
				[]int{30, 23, 19},
				[]int{84, 65, 58},
				[]int{138, 107, 97},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SquareMultiply(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SquareMultiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			"2x2 matrix",
			args{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
				[][]int{
					[]int{2, 0},
					[]int{1, 2},
				},
			},
			[][]int{
				[]int{4, 4},
				[]int{10, 8},
			},
			false,
		},
		{
			"2x2 matrix2",
			args{[][]int{
				[]int{1, 3},
				[]int{7, 5},
			},
				[][]int{
					[]int{6, 8},
					[]int{4, 2},
				},
			},
			[][]int{
				[]int{18, 14},
				[]int{62, 66},
			},
			false,
		},
		{
			"3x3 matrix",
			args{[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
				[][]int{
					[]int{9, 7, 8},
					[]int{6, 5, 4},
					[]int{3, 2, 1},
				},
			},
			[][]int{
				[]int{30, 23, 19},
				[]int{84, 65, 58},
				[]int{138, 107, 97},
			},
			false,
		},
		{
			"4x6 * 6x4",
			args{[][]int{
				[]int{1, 2, 3, 4, 5, 6},
				[]int{7, 8, 9, 10, 11, 12},
				[]int{13, 14, 15, 16, 17, 18},
				[]int{19, 20, 21, 22, 23, 24},
			}, [][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
				[]int{17, 18, 19, 20},
				[]int{21, 22, 23, 24},
			}},
			[][]int{
				[]int{301, 322, 343, 364},
				[]int{697, 754, 811, 868},
				[]int{1093, 1186, 1279, 1372},
				[]int{1489, 1618, 1747, 1876},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Multiply(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SquareMultiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareMultiplyStrassen(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			"2x2 matrix",
			args{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
				[][]int{
					[]int{2, 0},
					[]int{1, 2},
				},
			},
			[][]int{
				[]int{4, 4},
				[]int{10, 8},
			},
			false,
		},
		{
			"2x2 matrix2",
			args{[][]int{
				[]int{1, 3},
				[]int{7, 5},
			},
				[][]int{
					[]int{6, 8},
					[]int{4, 2},
				},
			},
			[][]int{
				[]int{18, 14},
				[]int{62, 66},
			},
			false,
		},
		{
			"3x3 matrix",
			args{[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
				[][]int{
					[]int{9, 7, 8},
					[]int{6, 5, 4},
					[]int{3, 2, 1},
				},
			},
			[][]int{
				[]int{30, 23, 19},
				[]int{84, 65, 58},
				[]int{138, 107, 97},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SquareMultiplyStrassen(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SquareMultiplyStrassen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareMultiplyStrassen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateMatricies(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"invalid matricies",
			args{nil, nil},
			true,
		},
		{
			"nonsquare matricies1",
			args{[][]int{
				[]int{1, 1},
			},
				[][]int{
					[]int{1, 2, 3},
					[]int{1},
				},
			},
			true,
		},
		{
			"nonsquare matricies",
			args{[][]int{
				[]int{1, 1},
				[]int{1, 1},
			},
				[][]int{
					[]int{1, 2, 3},
					[]int{1},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateMatricies(tt.args.a, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("validateMatricies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_submatix(t *testing.T) {
	type args struct {
		input [][]int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"2x2",
			args{[][]int{
				[]int{18, 14},
				[]int{62, 66}},
				1, 1},
			[][]int{
				[]int{18},
			},
		},
		{
			"2x2",
			args{[][]int{
				[]int{18, 14},
				[]int{62, 66}},
				1, 2},
			[][]int{
				[]int{14},
			},
		},
		{
			"2x2",
			args{[][]int{
				[]int{18, 14},
				[]int{62, 66}},
				2, 1},
			[][]int{
				[]int{62},
			},
		},
		{
			"2x2",
			args{[][]int{
				[]int{18, 14},
				[]int{62, 66}},
				2, 2},
			[][]int{
				[]int{66},
			},
		},
		{
			"4x4",
			args{[][]int{
				[]int{18, 14, 2, 2},
				[]int{62, 66, 3, 3},
				[]int{62, 66, 4, 4},
				[]int{62, 66, 5, 5}},
				1, 1},
			[][]int{
				[]int{18, 14},
				[]int{62, 66},
			},
		},
		{
			"4x4",
			args{[][]int{
				[]int{18, 14, 2, 2},
				[]int{62, 66, 3, 3},
				[]int{62, 66, 4, 4},
				[]int{62, 66, 5, 5}},
				1, 2},
			[][]int{
				[]int{2, 2},
				[]int{3, 3},
			},
		},
		{
			"4x4",
			args{[][]int{
				[]int{18, 14, 2, 2},
				[]int{62, 66, 3, 3},
				[]int{62, 66, 4, 4},
				[]int{62, 66, 5, 5}},
				2, 1},
			[][]int{
				[]int{62, 66},
				[]int{62, 66},
			},
		},
		{
			"4x4",
			args{[][]int{
				[]int{18, 14, 2, 2},
				[]int{62, 66, 3, 3},
				[]int{62, 66, 4, 4},
				[]int{62, 66, 5, 5}},
				2, 2},
			[][]int{
				[]int{4, 4},
				[]int{5, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := submatix(tt.args.input, tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submatix() = %v, want %v", got, tt.want)
			}
		})
	}
}
