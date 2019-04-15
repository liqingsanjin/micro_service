package userservice

import "errors"

var (
	ErrUserNamePasswordEmpty   = errors.New("empty username or password")
	ErrWrongUserNameOrPassword = errors.New("wrong username or password")
)
