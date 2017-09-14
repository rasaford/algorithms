package helper

import (
	"math"
	"reflect"
	"testing"
)

func TestRandBetween(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"0 .. 10",
			args{0, 10},
			false,
		},
		{
			"invalid range",
			args{12, 10},
			true,
		},
		{
			"negative numbers",
			args{-5, 5},
			true,
		},
	}
	for _, tt := range tests {
		// this test is non-deterministic and has a 2/1<<15 % chance of failing
		t.Run(tt.name, func(t *testing.T) {
			min, max := math.MaxInt32, math.MinInt32
			for i := 0; i < 1<<15; i++ {
				rand, err := RandBetween(tt.args.a, tt.args.b)
				if (err != nil) != tt.wantErr {
					t.Errorf("RandBetween() failed on %d, %d with %v", tt.args.a, tt.args.b, err)
				} else {
					if rand > max {
						max = rand
					}
					if rand < min {
						min = rand
					}
				}
			}
			if !tt.wantErr && (min != tt.args.a || max != tt.args.b) {
				t.Errorf("RandBetween() did not produce values in the range [%d .. %d] ", tt.args.a, tt.args.b)
			}
		})
	}
}

func TestFindMinMax(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"-50..20",
			args{GenerateRandom(1<<10, -50, 20)},
			-50,
			20,
		},
		{
			"0..0",
			args{GenerateRandom(1<<10, 0, 0)},
			0,
			0,
		},
		{
			"empty array",
			args{GenerateRandom(0, 0, 222)},
			0,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max := FindMinMax(tt.args.array)
			if min != tt.want {
				t.Errorf("FindMinMax() got = %v, want %v", min, tt.want)
			}
			if max != tt.want1 {
				t.Errorf("FindMinMax() got1 = %v, want %v", max, tt.want1)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	a1 := math.MaxInt16
	b1 := 5
	a2 := a1
	b2 := b1
	a3 := math.MinInt32
	b3 := -1 << 20
	a4 := a3
	b4 := b3
	type args struct {
		a *int
		b *int
	}
	tests := []struct {
		name       string
		args, want args
	}{
		{
			"2 positive",
			args{&a1, &b1},
			args{&b2, &a2},
		},
		{
			"2 negative",
			args{&a3, &b3},
			args{&b4, &a4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Swap(tt.args.a, tt.args.b)
			if *tt.args.a != *tt.want.a || *tt.args.b != *tt.want.b {
				t.Errorf("Swap() did not swap the values correctly")
			}
		})
	}
}

func TestClone(t *testing.T) {
	array := GenerateRandom(1<<10, 0, 1234)
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"random elements",
			args{array},
			array,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
