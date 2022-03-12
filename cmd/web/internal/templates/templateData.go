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
	Snips        []models.Snip
	SelectedUser SelectedUserTemplateData
}

type SelectedUserTemplateData struct {
	Name  string
	Snips []models.Snip
}
