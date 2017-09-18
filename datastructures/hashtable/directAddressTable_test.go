package hashtable

import (
	"reflect"
	"testing"
)

func TestNewDirectAddressTable(t *testing.T) {
	type args struct {
		keyMin int
		keyMax int
	}
	tests := []struct {
		name string
		args args
		want *directAddress
	}{
		{
			"empty universe",
			args{0, 0},
			nil,
		},
		{
			"invalid universe",
			args{55, 5},
			nil,
		},
		{
			"valid universe",
			args{1 << 4, 1 << 10},
			&directAddress{
				make([]interface{}, (1<<10)-(1<<4)),
				1 << 4,
				1 << 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDirectAddressTable(tt.args.keyMin, tt.args.keyMax); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirectAddressTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directAddress_Insert(t *testing.T) {
	str := "this is a test"
	type args struct {
		key   int
		value interface{}
	}
	tests := []struct {
		name    string
		d       *directAddress
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"invalid Key",
			NewDirectAddressTable(5, 55),
			args{3, &str},
			nil,
			true,
		},
		{
			"valid Key",
			NewDirectAddressTable(2, 64),
			args{63, &str},
			"this is a test",
			false,
		},
		{
			"valid Key",
			NewDirectAddressTable(2, 64),
			args{2, &str},
			"this is a test",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Insert(tt.args.key, tt.args.value)
			if got, err := tt.d.Search(tt.args.key); (err != nil) != tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirectAddressTable.Insert() = %v, want %v err: %v", got, tt.want, err)
			}
		})
	}
}

func Test_directAddress_Search(t *testing.T) {
	table := NewDirectAddressTable(0, 64)
	table.Insert(5, "5")
	table.Insert(0, "0")
	table.Insert(63, "63")
	type args struct {
		key int
	}
	tests := []struct {
		name    string
		d       *directAddress
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"search invalid",
			table,
			args{-5},
			nil,
			true,
		},
		{
			"search invalid2",
			table,
			args{234234},
			nil,
			true,
		},
		{
			"search valid",
			table,
			args{5},
			"5",
			false,
		},
		{
			"search valid2",
			table,
			args{0},
			"0",
			false,
		},
		{
			"search valid3",
			table,
			args{63},
			"63",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Search(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("directAddress.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("directAddress.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directAddress_Delete(t *testing.T) {
	str := "this is a test"
	type args struct {
		key   int
		value interface{}
	}
	tests := []struct {
		name string
		d    *directAddress
		args args
		want interface{}
	}{
		{
			"invalid Key",
			NewDirectAddressTable(5, 55),
			args{3, &str},
			nil,
		},
		{
			"valid Key",
			NewDirectAddressTable(2, 64),
			args{63, &str},
			nil,
		},
		{
			"valid Key",
			NewDirectAddressTable(2, 64),
			args{2, &str},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Insert(tt.args.key, tt.args.value)
			tt.d.Delete(tt.args.key)
			if got, _ := tt.d.Search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirectAddressTable.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
