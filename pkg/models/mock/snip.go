package mock

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"time"
)

var mockSnip = models.Snip{
	Id:      1,
	User:    "rhodeon",
	Title:   "mock snip",
	Content: "this is a mock snip",
	Created: time.Time{},
}

type SnipController struct {
	models.SnipController
}

func (c *SnipController) Insert(user string, title string, content string) (int, error) {
	return 2, nil
}

func (c *SnipController) Get(id int) (models.Snip, error) {
	switch id {
	case 1:
		return mockSnip, nil
	default:
		return models.Snip{}, models.ErrNoRecord
	}
}

func (c *SnipController) Latest(limit int) ([]models.Snip, error) {
	return []models.Snip{mockSnip}, nil
}
