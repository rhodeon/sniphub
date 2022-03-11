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
	// check if id is in the range of mockUsers indices
	if (0 < id) && (id < len(mockUsers)-1) {
		return mockUsers[id], nil
	}
	return models.User{}, models.ErrInvalidUser
}

func (c *UserController) GetSnips(username string) ([]models.Snip, error) {
	return []models.Snip{mockSnip}, nil
}
