package templates

import (
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
)

type TemplateData struct {
	// default data
	CurrentYear     int
	CsrfToken       string
	FlashMessage    string
	IsAuthenticated bool
	User            models.User

	// data from submitted form
	Form *forms.Form

	// data from database
	Snip         models.Snip
	Home         HomePageData
	SelectedUser SelectedUserData
}

type HomePageData struct {
	Snips    []models.Snip
	Current  int
	Next     int
	Previous int
	Last     bool
}

type SelectedUserData struct {
	Name  string
	Snips []models.Snip
}

type SnipData struct {
	models.Snip
	IsAuthenticated bool
	IsAuthor        bool
	CsrfToken       string
}
