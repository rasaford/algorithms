package hashtable

import (
	"encoding/binary"
	"math/big"
	"math/bits"
	"runtime"
)

func hashDivision(key, max uint) uint {
	return key % max
}

func HashMultiply64(key, max uint64) uint64 {
	maxLen := bits.Len(uint(max))
	wordSize := 32
	hashConstant := uint64(2.654435769e+09)
	if runtime.GOARCH == "amd64" {
		wordSize = 64
		hashConstant = 1.1400714819323198485e+019
	}
	hashBig := big.NewInt(0).SetUint64(hashConstant)
	keyBig := big.NewInt(0).SetUint64(uint64(key))
	hashBig.Mul(hashBig, keyBig)
	bytes := hashBig.Bytes()
	start := len(bytes) - wordSize/8
	bytes = bytes[start:]
	return uint64(highOrderBits(bytes, maxLen-1))
}

func HashMultiply32(key, max uint32) uint32 {
	maxLen := bits.Len(uint(max))
	hashConstant := uint64(2.654435769e+09) // (sqrt(5) - 1) * 2^(wordSize - 1)
	res := hashConstant * uint64(key)
	res &= 0x00000000FFFFFFFF
	res = res >> uint(32-maxLen+1)
	return uint32(res)
}

// highOrderBits retuns the first n high order bits of the given []byte
// as an int64.
func highOrderBits(bytes []byte, n int) uint64 {
	wholeBytes := int(n / 8)
	remBits := uint(n % 8)
	rightShift := uint(8)
	first := bytes[:wholeBytes]
	if remBits != 0 {
		remainder := uint(bytes[wholeBytes]) & (0xFF << (8 - remBits))
		first = append(first, byte(remainder))
		rightShift = remBits
	}
	if len(first) != 8 {
		fill := make([]byte, 8-len(first))
		first = append(fill, first...)
	}
	return binary.BigEndian.Uint64(first) >> (8 - rightShift)
}
