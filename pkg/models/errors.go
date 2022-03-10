package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateUsername  = errors.New("models: duplicate username")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrInvalidUser        = errors.New("models: user not found")
)
