package main

import (
	"errors"
	"fmt"
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
	"github.com/rhodeon/sniphub/pkg/session"
	"net/http"
)

func (app *application) signupUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "signup.page.gohtml", &TemplateData{Form: forms.New(nil)})
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
	form.MatchesPattern(forms.Email, forms.EmailRX)
	form.MaxLength(255, forms.Username, forms.Email)
	form.MinLength(10, forms.Password)

	// reload page with existing errors
	if !form.Valid() {
		app.renderTemplate(
			w, r,
			"signup.page.gohtml",
			&TemplateData{
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
				&TemplateData{
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
				&TemplateData{
					Form: form,
				},
			)
			return
		} else {
			serverError(w, err)
			return
		}
	}

	// redirect to login page
	app.sessionManager.Put(r.Context(), session.KeyFlashMessage, session.RegistrationSuccessful)
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "login.page.gohtml", &TemplateData{Form: forms.New(nil)})
}

func (app *application) loginUserPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login user")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout user")
}
