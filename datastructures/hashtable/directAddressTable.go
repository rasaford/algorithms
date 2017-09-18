package hashtable

import (
	"fmt"
)

type directAddress struct {
	array      []interface{}
	uMin, uMax int
}

// NewDirectAddressTable creates a new array with indices in the range:
// keyMIn <= x < keyMax.
//
// It requires O(u) space with u := keyMax - keyMIn
func NewDirectAddressTable(keyMin, keyMax int) *directAddress {
	if keyMin >= keyMax {
		return nil
	}
	size := keyMax - keyMin
	return &directAddress{
		array: make([]interface{}, size),
		uMin:  keyMin,
		uMax:  keyMax,
	}
}

// Insert inserts a new key, value pair in the table if the key is in the range of
// valid keys.
//
// It runs in O(1) time
func (d *directAddress) Insert(key int, value interface{}) {
	if err := d.validateKey(key); err != nil {
		return
	}
	d.array[key-d.uMin] = value
}

// Search returns the stored value for the given key if it is in the range of
// valid keys.
//
// It runs in O(1) time
func (d *directAddress) Search(key int) (interface{}, error) {
	if err := d.validateKey(key); err != nil {
		return nil, err
	}
	return d.array[key-d.uMin], nil
}

// Delete removes a key form the table if it is in the range of valid keys.
//
// It runs in O(1) time
func (d *directAddress) Delete(key int) {
	if err := d.validateKey(key); err != nil {
		return
	}
	d.array[key-d.uMin] = nil
}

func (d *directAddress) validateKey(key int) error {
	if key > d.uMax || key < d.uMin {
		return fmt.Errorf("key %d is not inside the valid key range of %d<=x<=%d", key, d.uMin, d.uMax)
	}
	return nil
}
