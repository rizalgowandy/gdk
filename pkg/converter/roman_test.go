package converter

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "MMII",
			args: args{
				number: 2002,
			},
			want: "MMII",
		},
		{
			name: "MCMXCVIII",
			args: args{
				number: 1998,
			},
			want: "MCMXCVIII",
		},
		{
			name: "M",
			args: args{
				number: 1000,
			},
			want: "M",
		},
		{
			name: "CM",
			args: args{
				number: 900,
			},
			want: "CM",
		},
		{
			name: "D",
			args: args{
				number: 500,
			},
			want: "D",
		},
		{
			name: "CD",
			args: args{
				number: 400,
			},
			want: "CD",
		},
		{
			name: "C",
			args: args{
				number: 100,
			},
			want: "C",
		},
		{
			name: "XC",
			args: args{
				number: 90,
			},
			want: "XC",
		},
		{
			name: "L",
			args: args{
				number: 50,
			},
			want: "L",
		},
		{
			name: "XL",
			args: args{
				number: 40,
			},
			want: "XL",
		},
		{
			name: "X",
			args: args{
				number: 10,
			},
			want: "X",
		},
		{
			name: "IX",
			args: args{
				number: 9,
			},
			want: "IX",
		},
		{
			name: "V",
			args: args{
				number: 5,
			},
			want: "V",
		},
		{
			name: "IV",
			args: args{
				number: 4,
			},
			want: "IV",
		},
		{
			name: "I",
			args: args{
				number: 1,
			},
			want: "I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRoman(tt.args.number); got != tt.want {
				t.Errorf("ToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
