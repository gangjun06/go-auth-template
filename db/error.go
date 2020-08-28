package db

import "errors"

var (
	// ErrUserAlreadyVerified returns error
	ErrUserAlreadyVerified = errors.New("User already verified")
	// ErrItemNotFound returns error
	ErrItemNotFound = errors.New("Cannot Find Item")
)
