package main

import (
	"net/http"

	"github.com/justinas/alice"
)

const (
	homeRoute        = "/"
	showSnipRoute    = "/snip/"
	createSnipRoute  = "/snip/create/"
	staticRoute      = "/static/"
	latestSnipsRoute = "/snip/latest/"
)

func (app *application) routesHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, app.home)
	mux.HandleFunc(showSnipRoute, app.showSnip)
	mux.HandleFunc(createSnipRoute, app.createSnip)
	mux.HandleFunc(staticRoute, app.serveStaticFiles)
	mux.HandleFunc(latestSnipsRoute, app.showLatestSnips)

	middlewareChain := alice.New(recoverPanic, logRequests, secureHeaders)
	return middlewareChain.Then(mux)
}
