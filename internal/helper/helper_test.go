package helper

import (
	"math"
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
			"0..20",
			args{GenerateRandom(1<<10, 0, 20)},
			0,
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
			got, got1 := FindMinMax(tt.args.array)
			if got != tt.want {
				t.Errorf("FindMinMax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindMinMax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
