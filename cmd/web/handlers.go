package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rhodeon/sniphub/pkg/models"
)

// default home response
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.renderTemplate(w, r, "home.page.gohtml", nil)
}

// displays a specified snippet
func (app *application) showSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

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

	// show specified snip
	snipTemplate := &TemplateData{Snip: snip}
	app.renderTemplate(w, r, "show.page.gohtml", snipTemplate)
}

// allows user to create a snippet
func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, fmt.Sprintf("%s?id=%d", showSnipRoute, id), http.StatusSeeOther)
}

// serves static files
func (app *application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix(staticRoute, fileServer).ServeHTTP(w, r)
}

// displays latest snips
func (app *application) showLatestSnips(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		notFoundError(w)
		return
	}

	// set a default limit of 10 if the limit query is less than 1
	if limit < 1 {
		limit = 10
	}

	// fetch latest snips from the database
	snips, err := app.snips.Latest(limit)
	if err != nil {
		serverError(w, err)
		return
	}

	// display list of latest snips
	snipTemplate := &TemplateData{Snips: snips}
	app.renderTemplate(w, r, "latest.page.gohtml", snipTemplate)
}
