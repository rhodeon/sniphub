package mock

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"time"
)

var mockSnips = []models.Snip{
	{
		Id:      1,
		User:    "rhodeon",
		Title:   "mock snip",
		Content: "this is a mock snip",
		Created: time.Time{},
	},
	{
		Id:      2,
		User:    "rhodeon",
		Title:   "another mock snip",
		Content: "this is another mock snip",
		Created: time.Time{},
	},
}

type SnipController struct {
	models.SnipController
}

func (c *SnipController) Insert(user string, title string, content string) (int, error) {
	return len(mockSnips) + 1, nil
}

func (c *SnipController) Get(id int) (models.Snip, error) {
	// iterate over mockSnips to return snip with matching id
	for _, snip := range mockSnips {
		if snip.Id == id {
			return snip, nil
		}
	}
	return models.Snip{}, models.ErrNoRecord
}

func (c *SnipController) Latest() ([]models.Snip, error) {
	return mockSnips, nil
}
