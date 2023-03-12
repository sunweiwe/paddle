package errors

import "errors"

var (
	ErrFailedToGetUser = errors.New("cannot get user from context")
)
