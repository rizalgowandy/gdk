package echo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/peractio/gdk/pkg/converter"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultResponse(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{
			name:   "Success",
			expect: converter.ToStr(http.StatusInternalServerError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			res := GetDefaultResponse(c)
			assert.Equal(t, tt.expect, res.Code)
		})
	}
}

func TestGetSuccessResponse(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{
			name:   "Success",
			expect: converter.ToStr(http.StatusOK),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			res := GetSuccessResponse(c, 0, "")
			assert.Equal(t, tt.expect, res.Code)
		})
	}
}

func TestPostDefaultResponse(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{
			name:   "Success",
			expect: converter.ToStr(http.StatusInternalServerError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			res := PostDefaultResponse(c)
			assert.Equal(t, tt.expect, res.Code)
		})
	}
}
