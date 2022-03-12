package main

import (
	"errors"
	"fmt"
	session2 "github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
)

// Default home response
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.gohtml", nil)
}

// Displays a specified snippet
func (app *application) showSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 0 {
		notFoundError(w)
		return
	}

	// attempt to retrieve the snip
	snip, err := app.snips.Get(id)
	if err != nil {
		// return a 404 error if the id matches none in the database
		if errors.Is(err, models.ErrNoRecord) {
			notFoundError(w)
		} else {
			serverError(w, err)
		}
		return
	}

	// show confirmation flash and specified snip
	snipTemplate := &templates.TemplateData{Snip: snip}
	app.renderTemplate(w, r, "show.page.gohtml", snipTemplate)
}

// Displays snip creation form
func (app *application) createSnipGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "create.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// Creates snip from submitted form and
// Redirect user to view the newly created snip.
// Returns to the creation form on error.
func (app *application) createSnipPost(w http.ResponseWriter, r *http.Request) {
	// verify form's content
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// validate title and content
	form := forms.New(r.PostForm)
	form.Required(forms.Title, forms.Content)
	form.MaxLength(100, forms.Title)

	// redirect to the creation form on error
	if !form.Valid() {
		app.renderTemplate(
			w, r,
			"create.page.gohtml",
			&templates.TemplateData{
				Form: form,
			},
		)
		return
	}

	// save the snip in the database
	id, err := app.snips.Insert(
		app.getUserFromContext(r).Username,
		form.Values.Get(forms.Title),
		form.Values.Get(forms.Content),
	)
	if err != nil {
		serverError(w, err)
		return
	}

	// redirect user to view newly created snip
	app.sessionManager.Put(r.Context(), session2.KeyFlashMessage, session2.SnipCreated)
	http.Redirect(w, r, fmt.Sprintf("%s/%d", showSnipRoute, id), http.StatusSeeOther)
}

// Serves static files
func (app *application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}

// Displays latest snips.
// Default limit of 10 if the limit query is less than 1 or nonexistent/malformed.
func (app *application) showLatestSnips(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10 // default limit
	}

	// fetch latest snips from the database
	snips, err := app.snips.Latest(limit)
	if err != nil {
		serverError(w, err)
		return
	}

	// display list of the latest snips
	snipTemplate := &templates.TemplateData{Snips: snips}
	app.renderTemplate(w, r, "latest.page.gohtml", snipTemplate)
}

// showUserSnips displays the snips for a user.
func (app *application) showUserSnips(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	// retrieve user snips
	snips, err := app.users.GetSnips(username)
	if err != nil {
		serverError(w, err)
		return
	}

	user := templates.SelectedUserTemplateData{
		Name:  username,
		Snips: snips,
	}
	td := &templates.TemplateData{SelectedUser: user}
	app.renderTemplate(w, r, "user_snips.page.gohtml", td)
}
