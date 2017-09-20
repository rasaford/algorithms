package hashtable

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

func hashDivision(key, max uint) uint {
	return key % max
}

func hashDivision2(key, max uint) uint {
	res := key % max
	if res%2 == 0 {
		return res - 1
	}
	return res
}

// HashMultiply64 hashes the key to a value in the range o <= x <= max.
// The hashing is based on applying the following function
//
//	h(k, m) = floor(m * (k * A mod 1))
// A = (sqrt(5) -1) / 2 = s / 2^w
//
// The first w significant bits are used as the hash value
// Hashing takes O(1) time.
func HashMultiply64(key, max uint64) uint64 {
	maxLen := bits.Len(uint(max))
	hash := big.NewInt(0).SetUint64(1.1400714819323198485e+019) // (sqrt(5) - 1) * 2^(wordSize - 1)
	keyBig := big.NewInt(0).SetUint64(uint64(key))
	hash.Mul(hash, keyBig)
	hashB := hash.Bytes()
	wordSize := 64
	start := len(hashB) - wordSize/8
	hashB = hashB[start:] // bit masking for the lower half of the result
	return uint64(highOrderBits(hashB, maxLen-1))
}

// HashMultiply32 hashes the key to a value in the range o <= x <= max.
// The hashing is based on applying the following function
//
//	h(k, m) = floor(m * (k * A mod 1))
// A = (sqrt(5) -1) / 2 = s / 2^w
//
// The first w significant bits are used as the hash value
// Hashing takes O(1) time.
func HashMultiply32(key, max uint32) uint32 {
	maxLen := bits.Len(uint(max))
	hashConstant := uint64(2.654435769e+09) // (sqrt(5) - 1) * 2^(wordSize - 1)
	res := hashConstant * uint64(key)
	res &= 0x00000000FFFFFFFF // bit masking for the lower half of the result
	res >>= uint(32 - maxLen + 1)
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
