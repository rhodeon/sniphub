package mysql

import (
	"database/sql"
	"errors"

	"github.com/rhodeon/sniphub/pkg/models"
)

type SnipController struct {
	Db *sql.DB
}

// Inserts a new snip to the database.
func (c *SnipController) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO snips (title, content, created) 
	VALUES(?, ?, UTC_TIMESTAMP)`

	result, err := c.Db.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	// return last inserted id for future reference
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Fetches the snip with the specified id from the database.
func (c *SnipController) Get(id int) (*models.Snip, error) {
	stmt := `SELECT id, title, content FROM snips
	WHERE id = ?`

	row := c.Db.QueryRow(stmt, id)
	snip := &models.Snip{}

	// fetch and map data from database to snip instance
	err := row.Scan(&snip.Id, &snip.Title, &snip.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return snip, nil
}

// Fetches a list of latest snips from the database.
func (c *SnipController) Latest() ([]*models.Snip, error) {
	return nil, nil
}
