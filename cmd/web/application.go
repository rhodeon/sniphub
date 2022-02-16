package main

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/rhodeon/sniphub/pkg/models/mysql"
)

type application struct {
	snips          *mysql.SnipController
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}
