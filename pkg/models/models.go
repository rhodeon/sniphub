package models

import (
	"time"
)

type Snip struct {
	id      int
	title   string
	content string
	created time.Time
}
