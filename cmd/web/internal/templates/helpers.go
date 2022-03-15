package templates

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"html/template"
	"path/filepath"
	"strings"
	"time"
)

var templateFunctions = template.FuncMap{
	"formattedDate": formattedDate,
	"embedIntoSnip": embedIntoSnip,
}

// NewTemplateCache caches templates for rendering pages.
func NewTemplateCache(dir string) (map[string]*template.Template, error) {
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
		ts, err := template.New(name).Funcs(templateFunctions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// parse all layout template files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		// parse all component template files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.component.gohtml"))
		if err != nil {
			return nil, err
		}

		// save the template set for each page in the cache
		cache[name] = ts
	}

	return cache, err
}

// formattedDate formats time to be more readable.
func formattedDate(t time.Time) string {
	// return zero-time instances as empty strings
	if t.IsZero() {
		return ""
	}

	// convert time to UTC before formatting
	return t.UTC().Format("Jan 02, 2006 at 15:04")
}

// embedIntoSnip inserts the csrf token and authorship status which are needed
// in the snip data.
func embedIntoSnip(csrfToken string, username string, snip models.Snip) SnipData {
	snipData := SnipData{snip, false, false, csrfToken}

	// set authentication and authorship status
	if strings.TrimSpace(username) != "" {
		snipData.IsAuthenticated = true

		if username == snip.User {
			snipData.IsAuthor = true
		}
	}

	return snipData
}
