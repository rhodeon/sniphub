package mysql

import (
	"database/sql"
	"github.com/rhodeon/sniphub/pkg/models"
)

type UserController struct {
	Db *sql.DB
}

func (c *UserController) Insert() error {
	return nil
}

func (c *UserController) Authenticate() (int, error) {
	return 0, nil
}

func (c *UserController) Get() (*models.User, error) {
	return nil, nil
}
