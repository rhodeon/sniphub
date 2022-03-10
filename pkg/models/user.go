package models

import "time"

type User struct {
	Id             int
	Username       string
	Email          string // unique users_uc_email constraint
	HashedPassword string
	Created        time.Time
	Active         bool
}

type UserController interface {
	Insert(username string, email string, password string) error
	Authenticate(email string, password string) (int, error)
	Get(id int) (User, error)
	GetSnips(username string) ([]Snip, error)
}
