package main

import (
	"html/template"
	"path/filepath"

	"github.com/rhodeon/sniphub/pkg/models"
)

type TemplateData struct {
	CurrentYear int // to be displayed in the footer
	Snip        *models.Snip
	Snips       []*models.Snip
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
		// parse the page template file to a template set
		ts, err := template.ParseFiles(page)
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
		name := filepath.Base(page)
		cache[name] = ts
	}

	return cache, err
}
