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

// Authenticate verifies inputted email and password against those in the database.
// It returns the id of a user with valid credentials.
func (c *UserController) Authenticate(email string, password string) (int, error) {
	// retrieve id and hashed password at row with entered email
	stmt := `SELECT id, hashed_password FROM users WHERE email = ? AND active = TRUE`
	row := c.Db.QueryRow(stmt, email)

	var id int
	var hashedPassword []byte
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// wrong email
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// verify password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// wrong password
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

// Get retrieves a user details with the specified id.
func (c *UserController) Get(id int) (models.User, error) {
	// retrieve user details
	stmt := `SELECT id, username, email, created, active FROM users WHERE id = ?`
	row := c.Db.QueryRow(stmt, id)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created, &user.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// user not found
			return models.User{}, models.ErrInvalidUser
		} else {
			return models.User{}, err
		}
	}

	return *user, nil
}

// GetFromEmail retrieves the user with the specified email.
func (c *UserController) GetFromEmail(email string) (models.User, error) {
	// retrieve user details
	stmt := `SELECT id, username, email, created, active FROM users WHERE email = ?`
	row := c.Db.QueryRow(stmt, email)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created, &user.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// user not found
			return models.User{}, models.ErrInvalidUser
		} else {
			return models.User{}, err
		}
	}

	return *user, nil
}

// GetSnips retrieves the snips created by the specified user.
func (c *UserController) GetSnips(username string) ([]models.Snip, error) {
	stmt := `SELECT id, user, title, content, created FROM snips
	WHERE user = ?
	ORDER by created DESC`

	rows, err := c.Db.Query(stmt, username)
	if err != nil {
		return nil, err
	}
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

// ChangePassword compares the entered currentPassword against the hashed password of the user with
// the id and updates it to the new password if correct.
func (c *UserController) ChangePassword(id int, currentPassword string, newPassword string) error {
	// retrieve hashed password for comparison
	stmt := `SELECT hashed_password FROM users WHERE id = ?`
	row := c.Db.QueryRow(stmt, id)

	var hashedPassword []byte
	err := row.Scan(&hashedPassword)
	if err != nil {
		return err
	}

	// compare passwords
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(currentPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// wrong password
			return models.ErrInvalidCredentials
		} else {
			return err
		}
	}

	// update password in database
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return err
	}

	stmt = `UPDATE users SET hashed_password = ? WHERE id = ?`
	_, err = c.Db.Exec(stmt, newHashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}
