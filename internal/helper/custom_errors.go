package helper

import "errors"

var (
	ErrAlreadyExists = errors.New("resource already exists")
	ErrNotFound      = errors.New("resource not found")
	ErrInternal      = errors.New("internal server error")
)
