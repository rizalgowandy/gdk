package regex

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchStr(t *testing.T) {
	type args struct {
		pattern string
		input   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Invalid Empty Client Number",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "",
			},
			want: false,
		},
		{
			name: "Invalid Client Number TOFU-IS-LIFE",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "TOFU-IS-LIFE",
			},
			want: false,
		},
		{
			name: "Invalid Client Number 123456-TOFU-***-FOR-****-LIFE-1234",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "123456-TOFU-***-FOR-****-LIFE-1234",
			},
			want: false,
		},
		{
			name: "Invalid Client Number *******",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "*******",
			},
			want: false,
		},
		{
			name: "Invalid Client Number 1234567890",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "1234567890",
			},
			want: false,
		},
		{
			name: "Invalid Client Number *******1234567890",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "*******1234567890",
			},
			want: false,
		},
		{
			name: "Invalid Client Number 1234567890*******",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "1234567890*******",
			},
			want: false,
		},
		{
			name: "Valid Client Number 123456*1234",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "123456*1234",
			},
			want: true,
		},
		{
			name: "Valid Client Number 12345678***************12345678",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
				input:   "12345678***************12345678",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := MatchStr(tt.args.pattern, tt.args.input)
			assert.Equal(t, tt.wantErr, gotErr != nil)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{
			name: "New",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
			},
			wantNil: false,
			wantErr: false,
		},
		{
			name: "Already registered",
			args: args{
				pattern: `(\d){6}(\*)+(\d){4}`,
			},
			wantNil: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Register(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("Register() got = %v, wantNil %v", err, tt.wantNil)
				return
			}
		})
	}
}
