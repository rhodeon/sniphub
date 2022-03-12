package main

import (
	"errors"
	session2 "github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
	"net/http"
)

// signupUserGet displays the account registration form.
func (app *application) signupUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "signup.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// signupUserPost saves a new user account to the database.
func (app *application) signupUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// validate name, email and password
	form := forms.New(r.PostForm)
	form.Required(forms.Username, forms.Email, forms.Password)
	form.MaxLength(255, forms.Username, forms.Email)
	form.MatchesPattern(forms.Email, forms.EmailRX)
	form.MinLength(10, forms.Password)

	// reload page with existing errors
	if !form.Valid() {
		app.renderTemplate(
			w, r,
			"signup.page.gohtml",
			&templates.TemplateData{
				Form: form,
			},
		)
		return
	}

	err = app.users.Insert(
		form.Values.Get(forms.Username),
		form.Values.Get(forms.Email),
		form.Values.Get(forms.Password),
	)
	if err != nil {
		// check for duplicate username
		if errors.Is(err, models.ErrDuplicateUsername) {
			form.Errors.Add(forms.Username, "Username is already taken")
			app.renderTemplate(w, r,
				"signup.page.gohtml",
				&templates.TemplateData{
					Form: form,
				},
			)
			return
		}

		// check for duplicate email
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add(forms.Email, "Email already in use")
			app.renderTemplate(w, r,
				"signup.page.gohtml",
				&templates.TemplateData{
					Form: form,
				},
			)
		} else {
			serverError(w, err)
		}
		return
	}

	// redirect to login page
	app.sessionManager.Put(r.Context(), session2.KeyFlashMessage, session2.RegistrationSuccessful)
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

// loginUserGet displays the user login form.
func (app *application) loginUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "login.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// loginUserPost compares received email and password against the database,
// and redirects to the homepage with the user ID on success.
func (app *application) loginUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// verify email and password
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Values.Get(forms.Email), form.Values.Get(forms.Password))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add(forms.Generic, "Email or password is incorrect")
			app.renderTemplate(w, r,
				"login.page.gohtml",
				&templates.TemplateData{
					Form: form,
				},
			)
		} else {
			serverError(w, err)
		}
		return
	}

	// store id and redirect to homepage
	app.sessionManager.Put(r.Context(), session2.KeyUserId, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// logoutUser removes the user id session key, and redirects to the homepage.
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Remove(r.Context(), session2.KeyUserId)
	app.sessionManager.Put(r.Context(), session2.KeyFlashMessage, session2.LogoutSuccessful)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
