package server

import (
	"github.com/rhodeon/sniphub/pkg/mailer"
	"github.com/rhodeon/sniphub/pkg/models"
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	Config         Env
	TemplateCache  map[string]*template.Template
	SessionManager *scs.SessionManager
	Mailer         *mailer.Mailer

	// database controllers
	Snips models.SnipController
	Users models.UserController
}
