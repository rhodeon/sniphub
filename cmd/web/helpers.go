package main

import (
	"fmt"
	"net/http"
)

// Renders html template set on screen
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *TemplateData) {
	// retrieved the cached template set from the name
	ts, exists := app.templateCache[name]
	if !exists {
		serverError(w, fmt.Errorf("the template %s does not exist", name))
	}

	// execute the template set
	err := ts.Execute(w, td)
	if err != nil {
		serverError(w, err)
	}
}
