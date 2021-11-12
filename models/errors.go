package models

import "errors"

var (
	ErrTimeout = errors.New("timeout server error")
	ErrServer  = errors.New("internal server error")
)
