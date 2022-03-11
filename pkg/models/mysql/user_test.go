package mysql

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"testing"
	"time"
)

func TestUserController_Get(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	db, teardown := newTestDb(t)
	defer teardown()
	userController := UserController{Db: db}

	tests := []struct {
		name     string
		userId   int
		wantUser models.User
		wantErr  error
	}{
		{
			"valid id",
			1,
			models.User{
				Id:       1,
				Username: "rhodeon",
				Email:    "rhodeon@mail.com",
				Created:  time.Date(2022, 3, 8, 23, 0, 0, 0, time.UTC),
				Active:   true,
			},
			nil,
		},
		{
			"invalid id",
			8,
			models.User{},
			models.ErrInvalidUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := userController.Get(tt.userId)
			testhelpers.AssertError(t, gotErr, tt.wantErr)
			testhelpers.AssertStruct(t, gotUser, tt.wantUser)
		})
	}
}
