package mux

import (
	"net/http"
	"time"

	"github.com/peractio/gdk/pkg/converter"
)

// GetDefaultResponse is the default response for http get request
func GetDefaultResponse(r *http.Request) Response {
	return Response{
		Code:       converter.String(http.StatusInternalServerError),
		DisplayMsg: http.StatusText(http.StatusInternalServerError),
		RawMsg:     "default",
		RequestID:  r.Header.Get(HeaderXRequestID),
		ResultData: GetResultData{
			Param:         r.URL.Query(),
			GeneratedDate: time.Now().Format(time.RFC3339),
			TotalData:     "",
			Data:          nil,
		},
	}
}

// GetSuccessResponse is the success response for http get request
func GetSuccessResponse(r *http.Request, totalData int, data interface{}) Response {
	return Response{
		Code:       converter.String(http.StatusOK),
		DisplayMsg: http.StatusText(http.StatusOK),
		RawMsg:     "",
		RequestID:  r.Header.Get(HeaderXRequestID),
		ResultData: GetResultData{
			Param:         r.URL.Query(),
			GeneratedDate: time.Now().Format(time.RFC3339),
			TotalData:     converter.String(totalData),
			Data:          data,
		},
	}
}

// PostDefaultResponse is the default response for http post request
func PostDefaultResponse(r *http.Request) Response {
	return Response{
		Code:       converter.String(http.StatusInternalServerError),
		DisplayMsg: http.StatusText(http.StatusInternalServerError),
		RawMsg:     "default",
		RequestID:  r.Header.Get(HeaderXRequestID),
		ResultData: PostResultData{
			Param:        r.URL.Query(),
			ExecutedDate: time.Now().Format(time.RFC3339),
			RowsAffected: 0,
		},
	}
}

// PostSuccessResponse is the success response for http post request
func PostSuccessResponse(r *http.Request, param interface{}, rowsAffected int) Response {
	return Response{
		Code:       converter.String(http.StatusOK),
		DisplayMsg: http.StatusText(http.StatusOK),
		RawMsg:     "",
		RequestID:  r.Header.Get(HeaderXRequestID),
		ResultData: PostResultData{
			Param:        param,
			ExecutedDate: time.Now().Format(time.RFC3339),
			RowsAffected: rowsAffected,
		},
	}
}
