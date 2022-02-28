package main

import (
	"fmt"
	"github.com/rhodeon/sniphub/pkg/forms"
	"net/http"
)

func (app *application) signupUserGet(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "signup.page.gohtml", &TemplateData{Form: forms.New(nil)})
}

func (app *application) signupUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		clientError(w, http.StatusBadRequest)
	}

	// validate name, email and password
	form := forms.New(r.PostForm)
	form.Required(forms.Name, forms.Email, forms.Password)
	form.MatchesPattern(forms.Email, forms.EmailRX)
	form.MaxLength(255, forms.Name, forms.Email)
	form.MinLength(10, forms.Password)

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

	fmt.Fprintln(w, "Create a new user...")
}

func (app *application) loginUserGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login form")
}

func (app *application) loginUserPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login user")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout user")
}
