package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// default home response
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		serverError(w, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		serverError(w, err)
		return
	}
}

// shows an example snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		notFoundError(w)
		return
	}

	fmt.Fprintf(w, "Displaying snippet %d", id)
}

// allows user to create a snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// dummy data
	title := "Someone"
	content := "The man, the myth, the legend."

	id, err := app.snips.Insert(title, content)
	if err != nil {
		serverError(w, err)
		return
	}

	// redirect user to view newly created snip
	http.Redirect(w, r, fmt.Sprintf("%s?id=%d", showSnippetRoute, id), http.StatusSeeOther)
}

// serves static files
func (app *application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix(staticRoute, fileServer).ServeHTTP(w, r)
}
