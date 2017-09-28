package dynamic

import (
	"math/rand"
	"testing"
)

var prices = getPrices(500)

func TestCutRod(t *testing.T) {
	type args struct {
		prices []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"book example",
			args{prices, 6},
			132,
		},
		{
			"large example",
			args{prices, 27},
			594,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRod(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("CutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCutRodMem(t *testing.T) {
	type args struct {
		prices []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"book example",
			args{prices, 6},
			132,
		},
		{
			"large example",
			args{prices, 27},
			594,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRodMemoized(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("CutRodMemoized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCutRodBottomUp(t *testing.T) {
	type args struct {
		prices []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"book example",
			args{prices, 6},
			132,
		},
		{
			"large example",
			args{prices, 27},
			594,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRodBottomUp(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("CutRodBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCutRodPrint(t *testing.T) {
	type args struct {
		prices []int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"book example",
			args{prices, 6},
			132,
		},
		{
			"large example",
			args{prices, 27},
			594,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRodPrint(tt.args.prices, tt.args.length); got != tt.want {
				t.Errorf("CutRodBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
func getPrices(size int) []int {
	rand.Seed(23834894589)
	prices := make([]int, size+1)
	last := 1
	for i := range prices {
		last += rand.Intn(22)
		prices[i] = last
	}
	return prices
}
