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

// Fetches a list of the 10 latest snips from the database.
func (c *SnipController) Latest() ([]*models.Snip, error) {
	stmt := `SELECT id, title, content FROM snips
	ORDER by created
	DESC LIMIT 10`

	rows, err := c.Db.Query(stmt)
	if err != nil {
		return nil, err
	}

	// ensure close the rows are closed before the function is returned in case of an error
	defer rows.Close()

	// empty slice to hold 10 latest snips
	snips := []*models.Snip{}

	// populate snips slice with pointers of mapped snip data from the database
	for rows.Next() {
		snip := &models.Snip{}
		rows.Scan(&snip.Id, &snip.Title, &snip.Content)
		snips = append(snips, snip)
	}

	// return any error that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snips, nil
}
