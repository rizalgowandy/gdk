package errorx

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

// GetArrJSON returns the json byte of the converted array of errors.
//
// Example:
// [{"code":"internal","message":"Internal server error.","op":"userService.FindUserByID"}]
func GetArrJSON(input error) []byte {
	all := GetArr(input)
	if all == nil {
		return nil
	}

	res, err := json.Marshal(all)
	if err != nil {
		return nil
	}

	return res
}

// GetArr returns the error as the converted array of errors.
//
// Example:
// [
//
//	{
//	  "code": "internal",
//	  "message": "Internal server error.",
//	  "op": "userService.FindUserByID"
//	},
//	{
//	  "code": "gateway",
//	  "message": "Gateway server error.",
//	  "op": "accountGateway.FindUserByID"
//	},
//	{
//	  "message": "Unknown error.",
//	  "op": "io.Write"
//	}
//
// ]
func GetArr(input error) []Error {
	e, ok := input.(*Error)
	if !ok {
		return nil
	}

	sub, ok := e.Err.(*Error)
	if ok {
		sub = nil
	}

	res := []Error{
		{
			Code:    e.Code,
			Message: e.Message,
			Op:      e.Op,
			Err:     sub,
		},
	}

	if !ok && e.Err != nil {
		res = append(res, Error{
			Code:    standard,
			Message: e.Err.Error(),
		})
	}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, GetArr(subErr)...)
	return res
}
