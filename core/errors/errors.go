package errors

import "errors"

var (

	// Context
	ErrFailedToGetUser      = errors.New("cannot get user from context")
	ErrFailedToGetRequestID = errors.New("cannot get the requestID from context")
)
