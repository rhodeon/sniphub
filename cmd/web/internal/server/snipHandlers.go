package server

import (
	"errors"
	"github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
)

// showSnip a specified snippet
func (app *Application) showSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 0 {
		notFoundError(w)
		return
	}

	// attempt to retrieve the snip
	snip, err := app.Snips.Get(id)
	if err != nil {
		// return a 404 error if the id matches none in the database
		if errors.Is(err, models.ErrNoRecord) {
			notFoundError(w)
		} else {
			serverError(w, err)
		}
		return
	}

	// show specified snip
	snipTemplate := &templates.TemplateData{Snip: snip}
	app.renderTemplate(w, r, "show.page.gohtml", snipTemplate)
}

// createSnipGet displays snip creation form
func (app *Application) createSnipGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "create.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// createSnipPost creates a snip from a submitted form and
// redirects the client to view the newly created snip.
// Returns to the creation form on error.
func (app *Application) createSnipPost(w http.ResponseWriter, r *http.Request) {
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
	id, err := app.Snips.Insert(
		app.getUserFromContext(r).Username,
		form.Values.Get(forms.Title),
		form.Values.Get(forms.Content),
	)
	if err != nil {
		serverError(w, err)
		return
	}

	// redirect user to view newly created snip
	app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.SnipCreated)
	http.Redirect(w, r, path.Join(showSnipRoute, strconv.Itoa(id)), http.StatusSeeOther)
}

// showUserSnips displays the snips created by a user.
func (app *Application) showUserSnips(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	// retrieve user snips
	snips, err := app.Users.GetSnips(username)
	if err != nil {
		serverError(w, err)
		return
	}

	user := templates.SelectedUserData{
		Name:  username,
		Snips: snips,
	}
	td := &templates.TemplateData{SelectedUser: user}
	app.renderTemplate(w, r, "user.page.gohtml", td)
}

func (app *Application) editSnipGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 0 {
		notFoundError(w)
		return
	}

	// attempt to retrieve the snip
	snip, err := app.Snips.Get(id)
	if err != nil {
		// return a 404 error if the id matches none in the database
		if errors.Is(err, models.ErrNoRecord) {
			notFoundError(w)
		} else {
			serverError(w, err)
		}
		return
	}

	app.renderTemplate(w, r, "edit_snip.page.gohtml", &templates.TemplateData{
		Form: &forms.Form{
			Values: url.Values{
				forms.Title:   {snip.Title},
				forms.Content: {snip.Content},
			},
			Errors: nil,
		},
		Snip: snip,
	})
}

func (app *Application) editSnipPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 0 {
		notFoundError(w)
		return
	}

	// verify form's content
	err = r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// validate title and content
	form := forms.New(r.PostForm)
	form.Required(forms.Title, forms.Content)
	form.MaxLength(100, forms.Title)

	// redirect to the creation form on an error
	if !form.Valid() {
		app.renderTemplate(
			w, r,
			"edit_snip.page.gohtml",
			&templates.TemplateData{
				Form: form,
			},
		)
		return
	}

	// update snip in the database
	err = app.Snips.Update(
		id,
		form.Values.Get(forms.Title),
		form.Values.Get(forms.Content),
	)
	if err != nil {
		serverError(w, err)
		return
	}

	// redirect to view newly edited snip
	app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.SnipEdited)
	http.Redirect(w, r, path.Join(showSnipRoute, strconv.Itoa(id)), http.StatusSeeOther)
}
