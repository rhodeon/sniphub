package main

import (
	"bytes"
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

	// instantiate buffer to hold template in case an error occurs
	buf := new(bytes.Buffer)

	// execute the template set
	err := ts.Execute(buf, td)
	if err != nil {
		serverError(w, err)
		return
	}

	// without an error, transfer the buffer's data to the response writer
	buf.WriteTo(w)
}
