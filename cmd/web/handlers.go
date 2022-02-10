package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhodeon/sniphub/pkg/models"
)

// default home response
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.gohtml", nil)
}

// displays a specified snippet
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

	// show specified snip
	snipTemplate := &TemplateData{Snip: snip}
	app.renderTemplate(w, r, "show.page.gohtml", snipTemplate)
}

// allows user to create a snippet
func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	// dummy data
	title := "Someone"
	content := "The man, the myth, the legend."

	id, err := app.snips.Insert(title, content)
	if err != nil {
		serverError(w, err)
		return
	}

	// redirect user to view newly created snip
	http.Redirect(w, r, fmt.Sprintf("%s/%d", showSnipRoute, id), http.StatusSeeOther)
}

// serves static files
func (app *application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}

// displays latest snips
// default limit of 10 if the limit query is less than 1 or nonexistent/malformed
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

	// display list of latest snips
	snipTemplate := &TemplateData{Snips: snips}
	app.renderTemplate(w, r, "latest.page.gohtml", snipTemplate)
}
