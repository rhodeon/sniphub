package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateUsername  = errors.New("models: duplicate username")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Snip struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

type User struct {
	Id             int
	Username       string
	Email          string // unique users_uc_email constraint
	HashedPassword string
	Created        time.Time
	Active         bool
}
