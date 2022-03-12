package models

import (
	"time"
)

type Snip struct {
	Id      int
	User    string
	Title   string
	Content string
	Created time.Time
}

type SnipController interface {
	Insert(user string, title string, content string) (int, error)
	Get(int) (Snip, error)
	Latest() ([]Snip, error)
}
