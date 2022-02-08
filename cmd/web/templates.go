package main

import "github.com/rhodeon/sniphub/pkg/models"

type TemplateData struct {
	Snip *models.Snip
	Snips []*models.Snip
}