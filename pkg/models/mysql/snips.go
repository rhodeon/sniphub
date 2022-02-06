package mysql

import (
	"database/sql"

	"github.com/rhodeon/sniphub/pkg/models"
)

type SnipController struct {
	Db *sql.DB
}

func (c *SnipController) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO snips (title, content, created) 
	VALUES(?, ?, UTC_TIMESTAMP)`

	result, err := c.Db.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (c *SnipController) Get(id int) (*models.Snip, error) {
	return nil, nil
}

func (c *SnipController) Latest() ([]*models.Snip, error) {
	return nil, nil
}
