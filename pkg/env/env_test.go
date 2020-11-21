package env

import (
	"os"
	"reflect"
	"testing"
)

func TestGetCurrent(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  string
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Alpha)
			},
			want: Alpha,
		},
		{
			name: "Env is missing",
			input: func() {
				_ = os.Setenv("GDK_ENV", "")
			},
			want: Development,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := GetCurrent()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  bool
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Development)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := IsDevelopment()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  bool
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Alpha)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := IsAlpha()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIsBeta(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  bool
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Beta)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := IsBeta()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIsStaging(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  bool
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Staging)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := IsStaging()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name  string
		input func()
		want  bool
	}{
		{
			name: "Env is available",
			input: func() {
				_ = os.Setenv("GDK_ENV", Production)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			once.Reset()
			tt.input()
			got := IsProduction()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
