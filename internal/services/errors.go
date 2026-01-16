package services

import "errors"

var (
	//Sign up errors
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserCreationFailed = errors.New("could not create user")

	//Sign in errors
	ErrIncorrectUser   = errors.New("user or password incorrect")
	ErrUserNotVerified = errors.New("user was not verified")

	//Create account errors
	ErrAccountAlreadyExists = errors.New("account already exists")

	//Transaction errors
	ErrTransactionNotCorrectCurrency = errors.New("this account not accepts this type of currency")
	ErrCantUpdateBalance             = errors.New("can't update cash balance of account")
)
