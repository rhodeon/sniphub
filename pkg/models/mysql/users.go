package mysql

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/rhodeon/sniphub/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserController struct {
	Db *sql.DB
}

// Insert saves user credentials on account creation.
// The credentials include username, email and hashed password.
// The account creation date is also stored.
func (c *UserController) Insert(username string, email string, password string) error {
	// hash password to save
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (username, email, hashed_password, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP)`

	_, err = c.Db.Exec(stmt, username, email, hashedPassword)
	if err != nil {
		mySqlErr := &mysql.MySQLError{}
		if errors.As(err, &mySqlErr) {
			// check if the error is due to a duplicate email
			// i.e. the 'users_uc_email' constraint for unique emails is raised
			if mySqlErr.Number == errDuplicateEntry {
				if strings.Contains(mySqlErr.Message, ConstraintUniqueUserName) {
					return models.ErrDuplicateUsername
				}
				if strings.Contains(mySqlErr.Message, ConstraintUniqueEmail) {
					return models.ErrDuplicateEmail
				}
			}
		}
		return err
	}
	return nil
}

func (c *UserController) Authenticate() (int, error) {
	return 0, nil
}

func (c *UserController) Get() (*models.User, error) {
	return nil, nil
}
