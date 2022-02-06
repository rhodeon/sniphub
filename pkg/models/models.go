package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snip struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}
