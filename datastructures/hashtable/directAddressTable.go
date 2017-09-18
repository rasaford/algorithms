package hashtable

import (
	"fmt"
)

type directAddress struct {
	array      []interface{}
	uMin, uMax int
}

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

func (d *directAddress) Insert(key int, value interface{}) {
	if err := d.validateKey(key); err != nil {
		return
	}
	d.array[key-d.uMin] = value
}

func (d *directAddress) Search(key int) (interface{}, error) {
	if err := d.validateKey(key); err != nil {
		return nil, err
	}
	return d.array[key-d.uMin], nil
}

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
