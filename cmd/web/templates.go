package main

import (
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
	"html/template"
	"path/filepath"
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
	SelectedUser SelectedUserTemplate
}

type SelectedUserTemplate struct {
	Name  string
	Snips []models.Snip
}

var templateFunctions = template.FuncMap{
	"formattedDate": formattedDate,
}

// Caches templates for rendering pages.
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// get a slice of all page template files to work with
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	// process each page file
	for _, page := range pages {
		name := filepath.Base(page)

		// associate the template function map with the template set
		// parse the page template file to a template set
		ts, err := template.New(name).Funcs(templateFunctions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// parse all layout template files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		// parse all partial template files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}

		// save the template set for each page in the cache
		cache[name] = ts
	}

	return cache, err
}
