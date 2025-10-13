package services

import "errors"

var (
	//Sign up errors
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserCreationFailed = errors.New("could not create user")

	//Sign in errors
	ErrIncorrectUser   = errors.New("user or password incorrect")
	ErrUserNotVerified = errors.New("user was not verified")
)
