package server

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
	"time"
)

// Renders html template set on screen
func (app *Application) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData) {
	// retrieved the cached template set from the name
	ts, exists := app.TemplateCache[name]
	if !exists {
		serverError(w, fmt.Errorf("the template %s does not exist", name))
	}

	// instantiate buffer to hold template in case an error occurs
	buf := new(bytes.Buffer)

	// execute the template set
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		serverError(w, err)
		return
	}

	// without an error, transfer the buffer's data to the response writer
	buf.WriteTo(w)
}

// Inserts default data for templates
func (app *Application) addDefaultData(td *templates.TemplateData, r *http.Request) *templates.TemplateData {
	if td == nil {
		td = &templates.TemplateData{}
	}

	td.CurrentYear = time.Now().Year()
	td.FlashMessage = app.SessionManager.PopString(r.Context(), session.KeyFlashMessage)
	td.CsrfToken = nosurf.Token(r)
	td.IsAuthenticated = app.isAuthenticated(r)
	td.User = app.getUserFromContext(r)
	return td
}
