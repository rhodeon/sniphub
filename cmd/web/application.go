package main

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager

	// database controllers
	snips models.SnipController
	users models.UserController
}
