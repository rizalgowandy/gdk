package filepathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGoMod(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindProjectAbs()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindGoMod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got)
		})
	}
}
