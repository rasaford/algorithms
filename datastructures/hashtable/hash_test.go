package hashtable

import (
	"math/rand"
	"testing"
)

var res uint64
var res2 uint32

func benchmark_HashMultiply(max uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := uint64(rand.Intn(int(max-1))) + 1
		res = HashMultiply64(r, max)
	}
}

func benchmark_HashMultiplySimple(max uint32, b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := uint32(rand.Intn(int(max-1))) + 1
		res2 = HashMultiply32(r, max)
	}
}

func Benchmark_HashMultiply10(b *testing.B) { benchmark_HashMultiply(1<<10, b) }
func Benchmark_HashMultiply11(b *testing.B) { benchmark_HashMultiply(1<<11, b) }
func Benchmark_HashMultiply12(b *testing.B) { benchmark_HashMultiply(1<<12, b) }
func Benchmark_HashMultiply13(b *testing.B) { benchmark_HashMultiply(1<<13, b) }
func Benchmark_HashMultiply14(b *testing.B) { benchmark_HashMultiply(1<<14, b) }
func Benchmark_HashMultiply20(b *testing.B) { benchmark_HashMultiply(1<<20, b) }

func Benchmark_HashMultiplySimple10(b *testing.B) { benchmark_HashMultiplySimple(1<<10, b) }
func Benchmark_HashMultiplySimple11(b *testing.B) { benchmark_HashMultiplySimple(1<<11, b) }
func Benchmark_HashMultiplySimple12(b *testing.B) { benchmark_HashMultiplySimple(1<<12, b) }
func Benchmark_HashMultiplySimple13(b *testing.B) { benchmark_HashMultiplySimple(1<<13, b) }
func Benchmark_HashMultiplySimple14(b *testing.B) { benchmark_HashMultiplySimple(1<<14, b) }
func Benchmark_HashMultiplySimple20(b *testing.B) { benchmark_HashMultiplySimple(1<<20, b) }

func Test_hashDivision(t *testing.T) {
	type args struct {
		key       uint
		maxOutput uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"large",
			args{5, 1 << 32},
			1 << 32,
		},
		{
			"small",
			args{1, 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.key; i <= tt.args.maxOutput; i += i {
				if got := hashDivision(tt.args.key, tt.args.maxOutput); got > tt.want {
					t.Errorf("hashDivision() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_hashMultiply64(t *testing.T) {
	type args struct {
		key       uint64
		maxOutput uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"large",
			args{5, 1 << 63},
			1 << 63,
		},
		{
			"small",
			args{1, 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.key; i <= tt.args.maxOutput; i += i {
				if got := HashMultiply64(tt.args.key, tt.args.maxOutput); got > tt.want {
					t.Errorf("hashMultiply64() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_hashMultiply32(t *testing.T) {
	type args struct {
		key       uint32
		maxOutput uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"large",
			args{5, 1 << 31},
			1 << 31,
		},
		{
			"small",
			args{1, 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := tt.args.key; i <= tt.args.maxOutput; i += i {
				if got := HashMultiply32(tt.args.key, tt.args.maxOutput); got > tt.want {
					t.Errorf("hashMultiply32() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
