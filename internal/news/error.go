package news

import "errors"

var (
	ErrNewsNotFound         = errors.New("news not found")
	ErrIncorrectCredentials = errors.New("incorrect json data")
)
