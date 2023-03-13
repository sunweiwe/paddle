package rpcerror

import (
	"fmt"
	"net/http"
)

type ErrorCode string

type RPCError struct {
	HTTPCode     int       `json:"-"`
	ErrorCode    ErrorCode `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
}

var (
	InternalError = RPCError{
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "InternalError",
	}
	ParameterError = RPCError{
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "InvalidParam",
	}
)

func (e RPCError) WithErrorMessageFormat(format string, param ...interface{}) RPCError {
	return RPCError{
		HTTPCode:     e.HTTPCode,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: fmt.Sprintf(format, param...),
	}
}
