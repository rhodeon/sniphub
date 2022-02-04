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
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(showSnippetRoute, showSnippet)
	mux.HandleFunc(createSnippetRoute, createSnippet)
	mux.HandleFunc(staticRoute, serveStaticFiles)
	return mux
}
