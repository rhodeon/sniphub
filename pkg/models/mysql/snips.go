package mysql

import (
	"database/sql"

	"github.com/rhodeon/sniphub/pkg/models"
)

type SnipController struct {
	Db *sql.DB
}

func (c *SnipController) Insert(title string, content string) (int, error) {
	return 0, nil
}

func (c *SnipController) Get(id int) (*models.Snip, error) {
	return nil, nil
}

func (c *SnipController) Latest() ([]*models.Snip, error) {
	return nil, nil
}
