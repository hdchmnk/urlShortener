package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLNotExist = errors.New("url not exist")
)
