package main

import (
	"html/template"

	"github.com/rhodeon/sniphub/pkg/models/mysql"
)

type application struct {
	snips         *mysql.SnipController
	templateCache map[string]*template.Template
}
