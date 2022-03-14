package mysql

import (
	"database/sql"
	"errors"
	"github.com/rhodeon/sniphub/pkg/models"
)

type SnipController struct {
	models.SnipController
	Db *sql.DB
}

// Insert inserts a new snip to the database.
func (c *SnipController) Insert(user string, title string, content string) (int, error) {
	stmt := `INSERT INTO snips (user, title, content, created) 
	VALUES(?, ?, ?, UTC_TIMESTAMP)`

	result, err := c.Db.Exec(stmt, user, title, content)
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

// Get fetches the snip with the specified id from the database.
func (c *SnipController) Get(id int) (models.Snip, error) {
	stmt := `SELECT id, user, title, content, created FROM snips
	WHERE id = ?`

	row := c.Db.QueryRow(stmt, id)
	snip := &models.Snip{}

	// fetch and map data from database to snip instance
	err := row.Scan(&snip.Id, &snip.User, &snip.Title, &snip.Content, &snip.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Snip{}, models.ErrNoRecord
		}
		return models.Snip{}, err
	}

	return *snip, nil
}

// Latest fetches a list of the 10 latest snips from the database.
func (c *SnipController) Latest(page int) ([]models.Snip, error) {
	stmt := `SELECT id, user, title, content, created FROM snips
	ORDER by created DESC
	LIMIT 10 OFFSET ?`

	rows, err := c.Db.Query(stmt, page*10)
	if err != nil {
		return nil, err
	}

	// ensure close the rows are closed before the function is returned in case of an error
	defer rows.Close()

	// empty slice to hold 10 latest snips
	var snips []*models.Snip

	// populate snips slice with pointers of mapped snip data from the database
	for rows.Next() {
		snip := &models.Snip{}
		_ = rows.Scan(&snip.Id, &snip.User, &snip.Title, &snip.Content, &snip.Created)
		snips = append(snips, snip)
	}

	// return any error that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// create a list to hold the value of each snip pointer
	var snipValues []models.Snip
	for _, snip := range snips {
		snipValues = append(snipValues, *snip)
	}
	return snipValues, nil
}

// Count returns the number of created snips.
func (c *SnipController) Count() int {
	stmt := `SELECT COUNT(*) FROM snips`

	row := c.Db.QueryRow(stmt)
	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

// Update replaces the title and content of the snip with the given id.
func (c *SnipController) Update(id int, title string, content string) error {
	stmt := `UPDATE snips SET title = ?, content = ? WHERE id = ?`

	_, err := c.Db.Exec(stmt, title, content, id)
	if err != nil {
		return err
	}
	return nil
}

// Clone makes a copy of the snip with the given id,
// changing the author and the creation date.
func (c *SnipController) Clone(id int, user string) (int, error) {
	stmt := `INSERT INTO snips (user, title, content, created)
	SELECT ?, title, content, UTC_TIMESTAMP FROM snips
	WHERE id = ?`

	result, err := c.Db.Exec(stmt, user, id)
	if err != nil {
		return 0, err
	}

	clonedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(clonedId), nil
}
