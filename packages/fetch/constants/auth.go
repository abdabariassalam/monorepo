package constants

import "errors"

var (
	ErrInvalidToken  = errors.New("token is invalid")
	ErrExpiredToken  = errors.New("token has expired")
	ErrRoleNotAdmin  = errors.New("you are not an admin")
	ErrTokenRequired = "token is required"
	ErrBadToken      = "bad token"
	RoleAdmin        = "admin"
)
