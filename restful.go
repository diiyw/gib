package goutils

import "github.com/diiyw/goutils/errors"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOK() Response {
	return Response{
		Code:    200,
		Message: "ok",
		Data:    nil,
	}
}

func ResponseSuccess(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func ResponseError(message string) Response {
	return Response{
		Code:    300,
		Message: message,
		Data:    nil,
	}
}

func ResponseWithError(err *errors.Error) Response {
	return Response{
		Code:    err.Code(),
		Message: err.Error(),
		Data:    nil,
	}
}
