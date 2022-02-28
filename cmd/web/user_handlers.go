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
	fmt.Fprintln(w, "Signup user")
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
