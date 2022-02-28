package main

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/rhodeon/sniphub/pkg/models/mysql"
)

type application struct {
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager

	// database controllers
	snips *mysql.SnipController
	users *mysql.UserController
}
