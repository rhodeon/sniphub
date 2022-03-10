package mock

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"time"
)

var mockUser = models.User{
	Id:             1,
	Username:       "rhodeon",
	Email:          "rhodeon@mail.com",
	HashedPassword: "",
	Created:        time.Time{},
	Active:         false,
}

type UserController struct {
	models.UserController
}

func (c *UserController) Insert(username string, email string, password string) error {
	return nil
}

func (c *UserController) Authenticate(email string, password string) (int, error) {
	return 2, nil
}

func (c *UserController) Get(id int) (models.User, error) {
	return mockUser, nil
}

func (c *UserController) GetSnips(username string) ([]models.Snip, error) {
	return []models.Snip{mockSnip}, nil
}
