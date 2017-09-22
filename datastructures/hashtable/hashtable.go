package hashtable

import (
	"github.com/rasaford/algorithms/datastructures/list"
)

// Table is a common interface for all the Hashtable implementations it supports:
// Insert, Search and Delete.
type Table interface {
	Insert(string, interface{})
	Search(string) interface{}
	Delete(string)
}

type chaining struct {
	values  []*list.List
	len     int
	maxLoad float64
}

type pair struct {
	key   string
	value interface{}
}

// NewChaining returns a new HashTable that uses chaining to deal with hash collisions.
// If the average length of the chains is >= 3 items, the table gets doubled in size.
func NewChaining() Table {
	return &chaining{
		values:  make([]*list.List, 8),
		len:     0,
		maxLoad: 3,
	}
}

// Insert inserts the given value at the hash location of the key.
// If the current loadfactor is > than the specified maximum, the underlying table
// gets doubled in size.
//
// It runs in O(1) expected time. Wort case is O(n) with n := len(table)
func (t *chaining) Insert(key string, val interface{}) {
	if t.loadFactor() > t.maxLoad {
		t.tableDouble()
	}
	hash := t.hash(key)
	if t.values[hash] == nil {
		t.values[hash] = list.New()
	}
	t.values[hash].Insert(&pair{key, val})
	t.len++
}

// Search seaches for the given key in the table and retuns an items if
// one is found. Otherwise nil is returned.
//
// It runs in O(1) time.
func (t *chaining) Search(key string) interface{} {
	hash := t.hash(key)
	list := t.values[hash]
	if list == nil || list.Len == 0 {
		return nil
	}
	head := list.Start().Prev
	for head != list.End() {
		head = head.Next
		pair := head.Value.(*pair)
		if pair.key == key {
			return pair.value
		}
	}
	return nil
}

// Delete deletes the value at the given key from the table.
//
// It runs in O(1) time.
func (t *chaining) Delete(key string) {
	hash := t.hash(key)
	list := t.values[hash]
	if list == nil {
		return
	}
	first := list.Start().Prev
	for first != list.End() {
		first = first.Next
		pair := first.Value.(*pair)
		if pair.key == key {
			list.Delete(first)
			t.len--
			return
		}
	}
}

func (t *chaining) loadFactor() float64 {
	return float64(t.len) / float64(len(t.values))
}

// tableDouble copies all keys / values from the current table into a newone, twice the size.
// All keys are rehashed to retain the properties of the hashtable.
func (t *chaining) tableDouble() {
	oldVal := t.values
	size := len(oldVal)
	t.values = make([]*list.List, size*2)
	for i := 0; i < size; i++ {
		list := oldVal[i]
		if list == nil {
			continue
		}
		first := list.Start().Prev
		for first != list.End() {
			first = first.Next
			pair := first.Value.(*pair)
			k, v := pair.key, pair.value
			t.Insert(k, v)
		}
	}
}

func (t *chaining) hash(key string) uint32 {
	return HashMultiply32(stringToInt(key), uint32(len(t.values)-1))
}

func stringToInt(s string) uint32 {
	res, j := uint32(0), uint32(0)
	for i := uint32(0); i < uint32(len(s)); i++ {
		j = i % 32
		res += uint32(s[i]) << j
	}
	return res
}

type openAddressing struct {
	values []*deletablePair
	len    int
}
type deletablePair struct {
	pair
	deleted bool
}

// NewOpenAddressing returns a Hashtable which uses only a single array to store the values
// and deals with hash collisions by rehashing the current key to a different slot until
// a non empty one is found.
func NewOpenAddressing() Table {
	return &openAddressing{
		values: make([]*deletablePair, 8),
		len:    0,
	}
}

// Insert inserts the given value at the hash location of the key.
// If the current loadfactor is > than the specified maximum, the underlying table
// gets doubled in size.
//
// It runs in O(1) expected time. Wort case is O(n) with n := len(table)
func (t *openAddressing) Insert(key string, value interface{}) {
	if t.loadFactor() > 0.5 {
		t.tableDouble()
	}
	round := 0
	for round != len(t.values) {
		hash := t.hash(key, round)
		if t.values[hash] == nil || t.values[hash].deleted {
			t.values[hash] = &deletablePair{
				pair:    pair{key, value},
				deleted: false,
			}
			t.len++
			return
		}
		round++
	}
}

// Search seaches for the given key in the table and retuns an items if
// one is found. Otherwise nil is returned.
//
// It runs in O(1) time.
func (t *openAddressing) Search(key string) interface{} {
	round := 0
	for round != len(t.values) {
		hash := t.hash(key, round)
		slot := t.values[hash]
		if slot != nil && !slot.deleted && slot.key == key {
			return slot.value
		}
		round++
	}
	return nil
}

// Delete deletes the value at the given key from the table.
//
// It runs in O(1) time.
func (t *openAddressing) Delete(key string) {
	round := 0
	for round != len(t.values) {
		hash := t.hash(key, round)
		slot := t.values[hash]
		if slot != nil && slot.key == key {
			t.values[hash].deleted = true
			t.len--
			return
		}
		round++
	}
}

// hash on an openAddressing table uses the double hashing scheme to deal with hash collisions.
func (t *openAddressing) hash(key string, round int) uint32 {
	num := uint(stringToInt(key))
	max := uint(len(t.values) - 1)
	return uint32((hashDivision(num, max) + uint(round)*hashDivision2(num, max)) % max)
}

func (t *openAddressing) loadFactor() float64 {
	return float64(t.len) / float64(len(t.values))
}

// tableDouble copies all keys / values from the current table into a newone, twice the size.
// All keys are rehashed to retain the properties of the hashtable.
func (t *openAddressing) tableDouble() {
	oldVal := t.values
	size := len(oldVal)
	t.values = make([]*deletablePair, size*2)
	for i := 0; i < size; i++ {
		list := oldVal[i]
		if list == nil || list.deleted {
			continue
		}
		t.Insert(list.key, list.value)
	}
}
