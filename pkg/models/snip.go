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
	Latest(page int) ([]Snip, error)
	Count() int
	Update(id int, title string, content string) error
	Clone(id int, user string) (int, error)
}
