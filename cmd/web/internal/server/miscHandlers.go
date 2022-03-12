package server

import (
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
)

// home is the default home handler. It shows the latest 10 snips in descending order of date.
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	// fetch the latest snips from the database
	snips, err := app.Snips.Latest()
	if err != nil {
		serverError(w, err)
		return
	}

	// display latest snips
	snipTemplate := &templates.TemplateData{Snips: snips}
	app.renderTemplate(w, r, "home.page.gohtml", snipTemplate)
}

func (app *Application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}
