package userservice

import "errors"

var (
	ErrUserNamePasswordEmpty = errors.New("empty username or password")
)
