package converter

import (
	"testing"
)

func TestStr(t *testing.T) {
	type args struct {
		v interface{}
	}
	type Person struct {
		Name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil value",
			args: args{
				nil,
			},
			want: "",
		},
		{
			name: "string correct format",
			args: args{
				"123 abc",
			},
			want: "123 abc",
		},
		{
			name: "int correct format",
			args: args{
				123,
			},
			want: "123",
		},
		{
			name: "int32 correct format",
			args: args{
				int32(123),
			},
			want: "123",
		},
		{
			name: "int64 correct format",
			args: args{
				int64(123),
			},
			want: "123",
		},
		{
			name: "bool correct format",
			args: args{
				true,
			},
			want: "true",
		},
		{
			name: "float32 correct format",
			args: args{
				float32(123.456),
			},
			want: "123.456",
		},
		{
			name: "float64 correct format",
			args: args{
				123.456,
			},
			want: "123.456",
		},
		{
			name: "uint8 list correct format",
			args: args{
				[]uint8{1, 2, 3},
			},
			want: "",
		},
		{
			name: "default value with error",
			args: args{
				make(chan int),
			},
			want: "",
		},
		{
			name: "default value no error",
			args: args{
				Person{"abc"},
			},
			want: `{"Name":"abc"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.v); got != tt.want {
				t.Errorf("Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "type string, accepted format, return true",
			args: args{
				v: "true",
			},
			want: true,
		},
		{
			name: "type string, accepted format, return false",
			args: args{
				v: "f",
			},
			want: false,
		},
		{
			name: "type string, unaccepted format, return false",
			args: args{
				v: "abc",
			},
			want: false,
		},
		{
			name: "type int = 0, return false",
			args: args{
				v: 0,
			},
			want: false,
		},
		{
			name: "type int != 0, return true",
			args: args{
				v: 1,
			},
			want: true,
		},
		{
			name: "default, return false",
			args: args{
				v: []byte("asd"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.v); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil value",
			args: args{
				nil,
			},
			want: 0,
		},
		{
			name: "string correct format",
			args: args{
				"123",
			},
			want: 123,
		},
		{
			name: "string bad format",
			args: args{
				"abc",
			},
			want: 0,
		},
		{
			name: "int correct format",
			args: args{
				123,
			},
			want: 123,
		},
		{
			name: "int32 correct format",
			args: args{
				int32(123),
			},
			want: 123,
		},
		{
			name: "int64 correct format",
			args: args{
				int64(123),
			},
			want: 123,
		},
		{
			name: "bool correct format",
			args: args{
				true,
			},
			want: 0,
		},
		{
			name: "float32 correct format",
			args: args{
				float32(123.00),
			},
			want: 123,
		},
		{
			name: "float64 correct format",
			args: args{
				123.456,
			},
			want: 123,
		},
		{
			name: "[]byte correct format",
			args: args{
				[]byte("123"),
			},
			want: 123,
		},
		{
			name: "[]byte incorrect format",
			args: args{
				[]byte("abc"),
			},
			want: 0,
		},
		{
			name: "default value with error",
			args: args{
				make(chan int),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.v); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "nil value",
			args: args{
				nil,
			},
			want: 0,
		},
		{
			name: "string correct format",
			args: args{
				"123",
			},
			want: 123,
		},
		{
			name: "string bad format",
			args: args{
				"abc",
			},
			want: 0,
		},
		{
			name: "int correct format",
			args: args{
				123,
			},
			want: 123,
		},
		{
			name: "int32 correct format",
			args: args{
				int32(123),
			},
			want: 123,
		},
		{
			name: "int64 correct format",
			args: args{
				int64(123),
			},
			want: 123,
		},
		{
			name: "bool correct format",
			args: args{
				true,
			},
			want: 0,
		},
		{
			name: "float32 correct format",
			args: args{
				float32(123.00),
			},
			want: 123,
		},
		{
			name: "float64 correct format",
			args: args{
				123.456,
			},
			want: 123,
		},
		{
			name: "[]byte correct format",
			args: args{
				[]byte("123"),
			},
			want: 123,
		},
		{
			name: "[]byte incorrect format",
			args: args{
				[]byte("abc"),
			},
			want: 0,
		},
		{
			name: "default value with error",
			args: args{
				make(chan int),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.v); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
