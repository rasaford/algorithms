package helper

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

// Swap uses pointers to two ints to swap them
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Clone clones the array and returns the copy
func Clone(input []int) []int {
	array := make([]int, len(input))
	copy(array, input)
	return array
}

// RandBetween generates random numbers in the range [a .. b] inclusive.
func RandBetween(a, b int) (int, error) {
	if a > b {
		return -1, fmt.Errorf("a:%d is larger than b:%d", a, b)
	}
	if a < 0 || b < 0 {
		return -1, fmt.Errorf("%f is less than 0", math.Min(float64(a), float64(b)))
	}
	return rand.Intn(b-a+1) + a, nil
}

// GenerateRandom generates a slice of random numbers between min and max inclusive
func GenerateRandom(length, min, max int) []int {
	offset := 0
	if min < 0 {
		offset = -min
		min = 0
		max += offset
	}
	out := make([]int, length)
	for i := range out {
		v, _ := RandBetween(min, max)
		out[i] = v - offset
	}
	return out
}

// FindMinMax retuns the minimum and maximum values in the given slice.
// It runs in O(3*n/2) time
func FindMinMax(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}
	min, max := array[0], array[0]
	if len(array)%2 == 0 {
		if array[0] < array[1] {
			min = array[0]
			max = array[1]
		} else {
			max = array[0]
			min = array[1]
		}
	}
	for i := 1; i < len(array); i += 2 {
		if array[i-1] > array[i] {
			if array[i-1] > max {
				max = array[i-1]
			}
			if array[i] < min {
				min = array[i]
			}
		} else {
			if array[i] > max {
				max = array[i]
			}
			if array[i-1] < min {
				min = array[i-1]
			}
		}
	}
	return min, max
}

func Print(node interface{}) {
	print(reflect.ValueOf(node), "", true, make(map[string]bool, 0))
}
func print(node reflect.Value, prefix string, tail bool, seen map[string]bool) {
	if !node.IsValid() {
		return
	}
	str := node.String()
	if res := seen[str]; res {
		return
	}
	seen[str] = true
	spacer1, spacer2 := "└── ", "    "
	if !tail {
		spacer1 = "├── "
		spacer2 = "│   "
	}
	fmt.Println(concat(prefix, spacer1, str, "\n"))
	w := concat(prefix, spacer2)
	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField()-1; i++ {
			print(v.Field(i), w, false, seen)
		}
		print(v.Field(v.NumField()-1), w, true, seen)
	}
}

func concat(strings ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}
