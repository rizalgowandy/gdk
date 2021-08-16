package nsqx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumerChain(t *testing.T) {
	type args struct {
		interceptors []ConsumerInterceptor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConsumerChain(tt.args.interceptors...)
			assert.NotNil(t, got)
		})
	}
}
