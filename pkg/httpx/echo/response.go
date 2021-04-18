package echo

import (
	"net/http"
	"time"

	"github.com/peractio/gdk/pkg/converter"

	"github.com/labstack/echo/v4"
)

// GetDefaultResponse is the default response for http get request
func GetDefaultResponse(c echo.Context) Response {
	return Response{
		Code:       converter.ToStr(http.StatusInternalServerError),
		DisplayMsg: http.StatusText(http.StatusInternalServerError),
		RawMsg:     "default",
		RequestID:  c.Response().Header().Get(echo.HeaderXRequestID),
		ResultData: GetResultData{
			Param:         c.QueryParams(),
			GeneratedDate: time.Now().Format(time.RFC3339),
			TotalData:     "",
			Data:          nil,
		},
	}
}

// GetSuccessResponse is the success response for http get request
func GetSuccessResponse(c echo.Context, totalData int, data interface{}) Response {
	return Response{
		Code:       converter.ToStr(http.StatusOK),
		DisplayMsg: http.StatusText(http.StatusOK),
		RawMsg:     "",
		RequestID:  c.Response().Header().Get(echo.HeaderXRequestID),
		ResultData: GetResultData{
			Param:         c.QueryParams(),
			GeneratedDate: time.Now().Format(time.RFC3339),
			TotalData:     converter.ToStr(totalData),
			Data:          data,
		},
	}
}

// PostDefaultResponse is the default response for http post request
func PostDefaultResponse(c echo.Context) Response {
	return Response{
		Code:       converter.ToStr(http.StatusInternalServerError),
		DisplayMsg: http.StatusText(http.StatusInternalServerError),
		RawMsg:     "default",
		RequestID:  c.Response().Header().Get(echo.HeaderXRequestID),
		ResultData: PostResultData{
			Param:        c.QueryParams(),
			ExecutedDate: time.Now().Format(time.RFC3339),
			RowsAffected: 0,
		},
	}
}

// PostSuccessResponse is the success response for http post request
func PostSuccessResponse(c echo.Context, param interface{}, rowsAffected int) Response {
	return Response{
		Code:       converter.ToStr(http.StatusOK),
		DisplayMsg: http.StatusText(http.StatusOK),
		RawMsg:     "",
		RequestID:  c.Response().Header().Get(echo.HeaderXRequestID),
		ResultData: PostResultData{
			Param:        param,
			ExecutedDate: time.Now().Format(time.RFC3339),
			RowsAffected: rowsAffected,
		},
	}
}
