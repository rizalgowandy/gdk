package password

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Password is empty",
			args:    args{},
			want:    "",
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				password: "Qwerty123!@#",
			},
			want:    "hash should not be empty",
			wantErr: false,
		},
	}

	enableLog := true
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != "") != (tt.want != "") {
				t.Errorf("Hash() got = %v, want %v", got, tt.want)
			}
			if enableLog {
				t.Logf("got = \n%#v\n\n", got)
			}
		})
	}
}

func TestIsMatch(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Password is empty",
			args: args{},
			want: false,
		},
		{
			name: "Hash is empty",
			args: args{},
			want: false,
		},
		{
			name: "Incorrect",
			args: args{
				password: "Qwerty123!@#",
				hash:     "123456",
			},
			want: false,
		},
		{
			name: "Correct",
			args: args{
				password: "Qwerty123!@#",
				hash:     "$2a$10$yra6SYVL/Q9IOQEjaw3ccOpbeh82mtewxeF6B.gRoOPJtL08dKpnq",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatch(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
