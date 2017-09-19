package prime

import (
	"testing"
)

var res []uint64

func benchmark_Sieve(max uint64, t *testing.B) {
	for i := 0; i < t.N; i++ {
		res = Sieve(max)
	}
}

func Benchmark_sieve32(b *testing.B)    { benchmark_Sieve(1<<5, b) }
func Benchmark_sieve64(b *testing.B)    { benchmark_Sieve(1<<6, b) }
func Benchmark_sieve128(b *testing.B)   { benchmark_Sieve(1<<7, b) }
func Benchmark_sieve256(b *testing.B)   { benchmark_Sieve(1<<8, b) }
func Benchmark_sieve512(b *testing.B)   { benchmark_Sieve(1<<9, b) }
func Benchmark_sieve1024(b *testing.B)  { benchmark_Sieve(1<<10, b) }
func Benchmark_sieve2048(b *testing.B)  { benchmark_Sieve(1<<11, b) }
func Benchmark_sieve4096(b *testing.B)  { benchmark_Sieve(1<<12, b) }
func Benchmark_sieve8192(b *testing.B)  { benchmark_Sieve(1<<13, b) }
func Benchmark_sieve16384(b *testing.B) { benchmark_Sieve(1<<14, b) }
func Benchmark_sieve20(b *testing.B)    { benchmark_Sieve(1<<20, b) }
