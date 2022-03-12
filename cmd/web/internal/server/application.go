package server

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	TemplateCache  map[string]*template.Template
	SessionManager *scs.SessionManager

	// database controllers
	Snips models.SnipController
	Users models.UserController
}