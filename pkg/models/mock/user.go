package mock

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"time"
)

var mockUsers = []models.User{
	{
		Id:             1,
		Username:       "rhodeon",
		Email:          "rhodeon@mail.com",
		HashedPassword: "",
		Created:        time.Time{},
		Active:         true,
	},
	{
		Id:             2,
		Username:       "crusoe",
		Email:          "crusoe@mail.com",
		HashedPassword: "",
		Created:        time.Time{},
		Active:         true,
	},
}

type UserController struct {
	models.UserController
}

func (c *UserController) Insert(username string, email string, password string) error {
	// check for duplicate username or email
	for _, user := range mockUsers {
		if username == user.Username {
			return models.ErrDuplicateUsername
		}
		if email == user.Email {
			return models.ErrDuplicateEmail
		}
	}
	return nil
}

func (c *UserController) Authenticate(email string, password string) (int, error) {
	return 2, nil
}

func (c *UserController) Get(id int) (models.User, error) {
	// iterate over mockUsers to return user with matching id
	for _, user := range mockUsers {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, models.ErrInvalidUser
}

func (c *UserController) GetSnips(username string) ([]models.Snip, error) {
	return mockSnips, nil
}
