package prime

import (
	"math"
	"runtime"
)

// func genPrime(target uint) uint {

// }

// func Primes(max uint64) []uint64 {
// 	cap := max / uint64(math.Log(float64(max)-1))
// 	all := make([]uint64, 0, max)
// 	segSize := uint64(math.Sqrt(float64(max)))
// 	cores := runtime.NumCPU()
// 	pool := sync.Pool{}
// 	pool.New = func() interface{} {
// 		return make([]bool, segSize)
// 	}
// 	basePrimes := SieveOfErastosthenes(segSize)
// 	all = append(all, basePrimes...)
// 	next := make(chan bool, 5)
// 	nextTurn := make([]chan bool, max/segSize+1)
// 	for i := uint64(0); i < max/segSize+1; i++ {
// 		nextTurn[i] = make(chan bool)
// 	}
// 	for segNum := uint64(1); segNum <= max/segSize; segNum++ {
// 		go fillSegments()
// 		next <- true
// 	}
// 	for i := 0; i < cores; i++ {
// 		next <- true
// 	}
// 	return all
// }

func SieveConcurrent(max uint64) []uint64 {
	cores := runtime.NumCPU()
	next := make(chan bool, cores)
	nums := make([]bool, max/2+1)
	m := sqrt(max)
	for i := uint64(3); i <= m; i += 2 {
		if !nums[i/2] {
			go func() {
				fill(nums, i, max)
				<-next
			}()
			next <- true
		}
	}
	for i := 0; i < cores; i++ {
		next <- true
	}
	ps := make([]uint64, 0, m)
	if max >= 2 {
		ps = append(ps, 2)
	}
	for i := uint64(3); i <= max; i += 2 {
		if !nums[i/2] {
			ps = append(ps, i)
		}
	}
	return ps
}

func fill(nums []bool, i, max uint64) {
	a := 3 * i
	for a <= max {
		nums[a/2] = true
		a += 2 * i
	}
}

func Sieve(max uint64) []uint64 {
	primes := make([]bool, max)
	output := make([]uint64, 0, max)
	for i := uint64(2); i < sqrt(max); i++ {
		if !primes[i] {
			for j := i * i; j < max; j += i {
				primes[j] = true

			}
		}
	}
	for i := 2; i < len(primes); i++ {
		if !primes[i] {
			output = append(output, uint64(i))
		}
	}
	return output
}

func SegmentedSieve(max uint64) {

}
func sqrt(input uint64) uint64 {
	return uint64(math.Sqrt(float64(input)))
}
