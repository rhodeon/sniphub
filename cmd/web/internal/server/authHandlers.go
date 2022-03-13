package server

import (
	"errors"
	"fmt"
	"github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/models"
	"net/http"
)

// signupUserGet displays the account registration form.
func (app *Application) signupUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "signup.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// signupUserPost saves a new user account to the database.
func (app *Application) signupUserPost(w http.ResponseWriter, r *http.Request) {
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

	err = app.Users.Insert(
		form.Values.Get(forms.Username),
		form.Values.Get(forms.Email),
		form.Values.Get(forms.Password),
	)
	if err != nil {
		// check for duplicate username
		if errors.Is(err, models.ErrDuplicateUsername) {
			form.Errors.Add(forms.Username, forms.ErrExistingUsername)
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
			form.Errors.Add(forms.Email, forms.ErrExistingEmail)
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
	app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.RegistrationSuccessful)
	http.Redirect(w, r, loginRoute, http.StatusSeeOther)
}

// loginUserGet displays the user login form.
func (app *Application) loginUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "login.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

// loginUserPost compares received email and password against the database,
// and redirects to the homepage with the user ID on success.
func (app *Application) loginUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// verify email and password
	form := forms.New(r.PostForm)
	id, err := app.Users.Authenticate(form.Values.Get(forms.Email), form.Values.Get(forms.Password))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add(forms.Generic, forms.ErrInvalidEmailOrPassword)
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
	app.SessionManager.Put(r.Context(), session.KeyUserId, id)
	http.Redirect(w, r, homeRoute, http.StatusSeeOther)
}

// logoutUser removes the user id session key, and redirects to the homepage.
func (app *Application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.SessionManager.Remove(r.Context(), session.KeyUserId)
	app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.LogoutSuccessful)
	http.Redirect(w, r, homeRoute, http.StatusSeeOther)
}

func (app *Application) changePasswordGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "password.page.gohtml", &templates.TemplateData{Form: forms.New(nil)})
}

func (app *Application) changePasswordPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	form := forms.New(r.PostForm)
	form.Required(forms.CurrentPassword, forms.NewPassword, forms.ConfirmPassword)
	form.MinLength(10, forms.NewPassword, forms.ConfirmPassword)
	form.AssertPasswordConfirmation(forms.NewPassword, forms.ConfirmPassword)

	for _, errs := range form.Errors {
		fmt.Printf("errs: %+v", errs)
	}

	//reload page with existing errors
	if !form.Valid() {
		fmt.Println("bhjbhhbjh")
		app.renderTemplate(
			w, r,
			"password.page.gohtml",
			&templates.TemplateData{
				Form: form,
			},
		)
		return
	}

	userId := app.getUserFromContext(r).Id
	err = app.Users.ChangePassword(
		userId,
		form.Values.Get(forms.CurrentPassword),
		form.Values.Get(forms.NewPassword),
	)

	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add(forms.CurrentPassword, forms.ErrIncorrectPassword)
			app.renderTemplate(
				w, r,
				"password.page.gohtml",
				&templates.TemplateData{
					Form: form,
				},
			)
		} else {
			serverError(w, err)
		}
		return
	}

	app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.PasswordChanged)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
