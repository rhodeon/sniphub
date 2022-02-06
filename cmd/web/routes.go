package main

import "net/http"

const (
	homeRoute          = "/"
	showSnippetRoute   = "/snippet/"
	createSnippetRoute = "/snippet/create/"
	staticRoute        = "/static/"
)

func (app *application) routesHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, app.home)
	mux.HandleFunc(showSnippetRoute, app.showSnippet)
	mux.HandleFunc(createSnippetRoute, app.createSnippet)
	mux.HandleFunc(staticRoute, app.serveStaticFiles)
	return mux
}
