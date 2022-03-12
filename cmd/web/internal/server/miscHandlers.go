package server

import "net/http"

// Default home response
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.gohtml", nil)
}

// Serves static files
func (app *Application) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}
