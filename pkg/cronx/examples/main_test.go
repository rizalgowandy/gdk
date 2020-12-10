package main

import (
	"testing"
)

func Test_payBill_Run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := payBill{}
			p.Run()
		})
	}
}

func Test_sendEmail_Run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := sendEmail{}
			e.Run()
		})
	}
}
