package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
	"strconv"
	"strings"
)

// home is the default home handler. It shows the latest 10 snips in descending order of date.
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	pageParam := chi.URLParam(r, "page")

	// ensure the root path has a page of 0
	if strings.TrimSpace(pageParam) == "" {
		pageParam = "0"
	}

	pageNumber, err := strconv.Atoi(pageParam)
	if err != nil || pageNumber < 0 {
		http.Redirect(w, r, homeRoute, http.StatusSeeOther)
		return
	}

	// fetch the latest snips from the database
	snips, err := app.Snips.Latest(pageNumber)
	if err != nil {
		serverError(w, err)
		return
	}

	// calculate total number of pages
	snipCount := app.Snips.Count()
	var pageCount int
	if snipCount != 0 && snipCount%10 == 0 {
		pageCount = snipCount/10 - 1
	} else {
		pageCount = snipCount / 10
	}

	// display latest snips
	snipTemplate := &templates.TemplateData{Home: templates.HomePageData{
		Snips:    snips,
		Current:  pageNumber,
		Next:     pageNumber + 1,
		Previous: pageNumber - 1,
		Last:     pageNumber == pageCount,
	}}
	app.renderTemplate(w, r, "home.page.gohtml", snipTemplate)
}

func (app *Application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}
