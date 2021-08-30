package httpx

import (
	"reflect"
	"testing"

	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
)

func TestDummy_Close(t *testing.T) {
	type fields struct {
		ReadErr  error
		CloseErr error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Error",
			fields: fields{
				ReadErr:  nil,
				CloseErr: errorx.E("error"),
			},
			wantErr: true,
		},
		{
			name:    "Success",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dummy{
				ReadErr:  tt.fields.ReadErr,
				CloseErr: tt.fields.CloseErr,
			}
			if err := d.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDummy_Read(t *testing.T) {
	type fields struct {
		ReadErr  error
		CloseErr error
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Error",
			fields: fields{
				ReadErr:  errorx.E("error"),
				CloseErr: nil,
			},
			args:    args{},
			wantN:   0,
			wantErr: true,
		},
		{
			name:    "Success",
			fields:  fields{},
			args:    args{},
			wantN:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dummy{
				ReadErr:  tt.fields.ReadErr,
				CloseErr: tt.fields.CloseErr,
			}
			gotN, err := d.Read(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Read() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewDummy(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *Dummy
	}{
		{
			name: "Success",
			args: args{},
			want: &Dummy{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDummy(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDummy() = %v, want %v", got, tt.want)
			}
		})
	}
}
