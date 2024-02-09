package admin

import "errors"

// ErrUser errors
var (
	ErrUserNotFound = errors.New("user not found")
)

// ErrNews errors
var (
	ErrNewsNotFound         = errors.New("news not found")
	ErrIncorrectCredentials = errors.New("incorrect json data")
)
