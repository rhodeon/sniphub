package main

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"github.com/rhodeon/sniphub/pkg/models"
	"net/http"
	"time"

	"github.com/rhodeon/sniphub/pkg/session"
)

// Renders html template set on screen
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData) {
	// retrieved the cached template set from the name
	ts, exists := app.templateCache[name]
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
func (app *application) addDefaultData(td *templates.TemplateData, r *http.Request) *templates.TemplateData {
	if td == nil {
		td = &templates.TemplateData{}
	}

	td.CurrentYear = time.Now().Year()
	td.FlashMessage = app.sessionManager.PopString(r.Context(), session.KeyFlashMessage)
	td.CsrfToken = nosurf.Token(r)
	td.IsAuthenticated = app.isAuthenticated(r)
	td.User = app.getUserFromContext(r)
	return td
}

// isAuthenticated returns true if the current request is from authenticated user,
// otherwise it returns false.
func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}

// getUserFromContext returns the current user data stored in a request's context.
// Returns an empty user struct on an error,
func (app *application) getUserFromContext(r *http.Request) models.User {
	user, ok := r.Context().Value(contextKeyUser).(models.User)
	if !ok {
		return models.User{}
	}
	return user
}
