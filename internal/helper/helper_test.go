package helper

import "testing"
import "math"

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
