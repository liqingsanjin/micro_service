package userservice

import "errors"

var (
	ErrUserNamePasswordEmpty   = errors.New("empty username or password")
	ErrWrongUserNameOrPassword = errors.New("wrong username or password")
	ErrUnAuthorized            = errors.New("unauthoriezd error")
	ErrReplyTypeInvalid        = errors.New("reply type invalid")
	ErrRequestTypeInvalid      = errors.New("request type invalid")
	ErrTokenContextMissing     = errors.New("token up for parsing was not passed through the context")
	ErrTokenInvalid            = errors.New("JWT Token was invalid")
	ErrTokenExpired            = errors.New("JWT Token is expired")
	ErrTokenMalformed          = errors.New("JWT Token is malformed")
	ErrTokenNotActive          = errors.New("token is not valid yet")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidParams           = errors.New("invalid params")
)
