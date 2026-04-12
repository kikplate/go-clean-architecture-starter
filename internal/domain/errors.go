package domain

import "errors"

var (
	ErrNotFound      = errors.New("not_found")
	ErrConflict      = errors.New("conflict")
	ErrInvalidInput  = errors.New("invalid_input")
)
