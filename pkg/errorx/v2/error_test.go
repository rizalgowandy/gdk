package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name  string
		input *Error
		want  string
	}{
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Err:     errors.New("standard-error"),
			},
			want: "standard-error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Error()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v\n"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestError_GetFields(t *testing.T) {
	type fields struct {
		Err          error
		Code         Code
		Fields       Fields
		OpTraces     []Op
		Message      Message
		Line         Line
		MetricStatus MetricStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   Fields
	}{
		{
			name: "Success",
			fields: fields{
				Err:  nil,
				Code: "",
				Fields: Fields{
					"K":  "V",
					"K2": "V2",
				},
				OpTraces:     nil,
				Message:      "",
				Line:         "",
				MetricStatus: "",
			},
			want: Fields{
				"K":  "V",
				"K2": "V2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Err:          tt.fields.Err,
				Code:         tt.fields.Code,
				Fields:       tt.fields.Fields,
				OpTraces:     tt.fields.OpTraces,
				Message:      tt.fields.Message,
				Line:         tt.fields.Line,
				MetricStatus: tt.fields.MetricStatus,
			}
			if got := e.GetFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
