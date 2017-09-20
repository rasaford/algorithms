package hashtable

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/rasaford/algorithms/datastructures/list"
)

func TestNewChaining(t *testing.T) {
	tests := []struct {
		name string
		want Table
	}{
		{
			"default table",
			&chaining{
				values:  make([]*list.List, 8),
				len:     0,
				maxLoad: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChaining(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChaining() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chaining_Insert(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	n := 1 << 12
	r := make([]args, n)
	for i := range r {
		r[i] = args{randString(20), rand.Intn(20)}
	}
	tests := []struct {
		name  string
		args  []args
		table Table
	}{
		{
			"single insert",
			[]args{
				args{"123456", 123456},
			},
			NewChaining(),
		},
		{
			"multiple insert",
			[]args{
				args{"123456", 123456},
				args{"123456345345q", 23942340},
				args{"alksdfj", 99324923},
				args{"1234562", 123456},
				args{"1234562345345q", 23942340},
				args{"alksd2fj", 99324923},
				args{"12345xx6", 123456},
				args{"123423423456", 123456},
				args{"12ölkdfjglöarkgj3456", 123456},
				args{"1öawiejfosiefj23456", 123456},
				args{"yaay hashing like a pro ;)", 123456},
			},
			NewChaining(),
		},
		{
			"multiple insert table doubling",
			r,
			NewChaining(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.table.Insert(arg.key, arg.val)
			}
			for _, arg := range tt.args {
				if res := tt.table.Search(arg.key); res == nil {
					t.Errorf("chaining.Insert() want: %s:%v got: %v", arg.key, arg.val, res)
				}
			}
		})
	}
}

func Test_chaining_Search(t *testing.T) {
	type search struct {
		key string
		val interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		search  []search
		args    args
		table   Table
		wantErr bool
	}{
		{
			"successful search",
			[]search{
				search{"123456", 123456},
			},
			args{"123456"},
			NewChaining(),
			false,
		},
		{
			"unsuccessful search",
			[]search{
				search{"123456", 123456},
			},
			args{"123"},
			NewChaining(),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.search {
				tt.table.Insert(arg.key, arg.val)
			}
			if res := tt.table.Search(tt.args.key); (res == nil) != tt.wantErr {
				t.Errorf("chaining.Insert() key: %s: got: %v", tt.args.key, res)
			}
		})
	}
}

func Test_chaining_Delete(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	n := 1 << 12
	r := make([]args, n)
	for i := range r {
		r[i] = args{randString(20), rand.Intn(20)}
	}
	tests := []struct {
		name  string
		args  []args
		table Table
	}{
		{
			"single insert",
			[]args{
				args{"123456", 123456},
			},
			NewChaining(),
		},
		{
			"multiple insert",
			[]args{
				args{"123456", 123456},
				args{"123456345345q", 23942340},
				args{"alksdfj", 99324923},
				args{"1234562", 123456},
				args{"1234562345345q", 23942340},
				args{"alksd2fj", 99324923},
				args{"12345xx6", 123456},
				args{"123423423456", 123456},
				args{"12ölkdfjglöarkgj3456", 123456},
				args{"1öawiejfosiefj23456", 123456},
				args{"yaay hashing like a pro ;)", 123456},
			},
			NewChaining(),
		},
		{
			"multiple insert table doubling",
			r,
			NewChaining(),
		},
		{
			"delete from empty list",
			[]args{},
			NewChaining(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.table.Insert(arg.key, arg.val)
			}
			for _, arg := range tt.args {
				tt.table.Delete(arg.key)
			}
			for _, arg := range tt.args {
				if res := tt.table.Search(arg.key); res != nil {
					t.Errorf("chaining.Insert() want: %s:%v got: %v", arg.key, arg.val, res)
				}
			}
		})
	}
}

func TestOpenAddressing(t *testing.T) {
	tests := []struct {
		name string
		want Table
	}{
		{
			"default table",
			&openAddressing{
				values: make([]*deletablePair, 8),
				len:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOpenAddressing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpenAddressing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openAddressing_Insert(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	n := 1 << 12
	r := make([]args, n)
	for i := range r {
		r[i] = args{randString(20), rand.Intn(20)}
	}
	tests := []struct {
		name  string
		args  []args
		table Table
	}{
		{
			"single insert",
			[]args{
				args{"123456", 123456},
			},
			NewOpenAddressing(),
		},
		{
			"multiple insert",
			[]args{
				args{"123456", 123456},
				args{"123456345345q", 23942340},
				args{"alksdfj", 99324923},
				args{"1234562", 123456},
				args{"1234562345345q", 23942340},
				args{"alksd2fj", 99324923},
				args{"12345xx6", 123456},
				args{"123423423456", 123456},
				args{"12ölkdfjglöarkgj3456", 123456},
				args{"1öawiejfosiefj23456", 123456},
				args{"yaay hashing like a pro ;)", 123456},
			},
			NewOpenAddressing(),
		},
		{
			"multiple insert table doubling",
			r,
			NewOpenAddressing(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.table.Insert(arg.key, arg.val)
			}
			for _, arg := range tt.args {
				if res := tt.table.Search(arg.key); res == nil {
					t.Errorf("chaining.Insert() want: %s:%v got: %v", arg.key, arg.val, res)
				}
			}
		})
	}
}

func Test_openAddressing_Search(t *testing.T) {
	type search struct {
		key string
		val interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		search  []search
		args    args
		table   Table
		wantErr bool
	}{
		{
			"successful search",
			[]search{
				search{"123456", 123456},
			},
			args{"123456"},
			NewOpenAddressing(),
			false,
		},
		{
			"unsuccessful search",
			[]search{
				search{"123456", 123456},
			},
			args{"123"},
			NewOpenAddressing(),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.search {
				tt.table.Insert(arg.key, arg.val)
			}
			if res := tt.table.Search(tt.args.key); (res == nil) != tt.wantErr {
				t.Errorf("chaining.Insert() key: %s: got: %v", tt.args.key, res)
			}
		})
	}
}

func Test_openAddressing_Delete(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	n := 1 << 10
	r := make([]args, n)
	for i := range r {
		r[i] = args{randString(40), rand.Intn(20)}
	}
	tests := []struct {
		name  string
		args  []args
		table Table
	}{
		{
			"single insert",
			[]args{
				args{"123456", 123456},
			},
			NewOpenAddressing(),
		},
		{
			"multiple insert",
			[]args{
				args{"123456", 123456},
				args{"123456345345q", 23942340},
				args{"alksdfj", 99324923},
				args{"1234562", 123456},
				args{"1234562345345q", 23942340},
				args{"alksd2fj", 99324923},
				args{"12345xx6", 123456},
				args{"123423423456", 123456},
				args{"12ölkdfjglöarkgj3456", 123456},
				args{"1öawiejfosiefj23456", 123456},
				args{"yaay hashing like a pro ;)", 123456},
			},
			NewOpenAddressing(),
		},
		{
			"multiple insert table doubling",
			r,
			NewOpenAddressing(),
		},
		{
			"delete from empty list",
			[]args{},
			NewOpenAddressing(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				tt.table.Insert(arg.key, arg.val)
			}
			for _, arg := range tt.args {
				tt.table.Delete(arg.key)
			}
			for _, arg := range tt.args {
				if res := tt.table.Search(arg.key); res != nil {
					t.Errorf("chaining.Insert() want: %s:%v got: %v", arg.key, arg.val, res)
				}
			}
		})
	}
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bit
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
