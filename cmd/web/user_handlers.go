package main

import (
	"fmt"
	"net/http"
)

func (app *application) signupUserGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sighup form")
}

func (app *application) signupUserPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sighup user")
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
