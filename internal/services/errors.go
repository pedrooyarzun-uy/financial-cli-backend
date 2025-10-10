package services

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserCreationFailed = errors.New("could not create user")
)
