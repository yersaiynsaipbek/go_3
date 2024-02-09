package auth

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrIncorrectCredentials = errors.New("incorrect username or password")
)
