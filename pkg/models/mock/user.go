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
		HashedPassword: "$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG",
		Created:        time.Time{},
		Active:         true,
	},
	{
		Id:             2,
		Username:       "crusoe",
		Email:          "crusoe@mail.com",
		HashedPassword: "BVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG$2a$12$NuTjWXm3KKntReFwyBVHy",
		Created:        time.Time{},
		Active:         true,
	},
}

// mock map to represent hashed passwords
var mockPasswordHashes = map[string]string{
	"qwerty123456": "$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG",
	"qwertyuiop":   "BVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG$2a$12$NuTjWXm3KKntReFwyBVHy",
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
	// check if representation of hashed password exists
	hashedPassword, exists := mockPasswordHashes[password]
	if !exists {
		return 0, models.ErrInvalidCredentials
	}

	// verify that the email and hashed password exist
	for _, user := range mockUsers {
		if email == user.Email && hashedPassword == user.HashedPassword {
			return user.Id, nil
		}
	}

	return 0, models.ErrInvalidCredentials
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
