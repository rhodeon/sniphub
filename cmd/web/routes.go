package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()
	router.Get(homeRoute, app.home)
	router.Get(showSnipRoute, app.showSnip)
	router.Post(createSnipRoute, app.createSnip)
	router.Get(staticRoute, app.serveStaticFiles)
	router.Get(latestSnipsRoute, app.showLatestSnips)

	middlewareChain := alice.New(recoverPanic, logRequests, secureHeaders)
	return middlewareChain.Then(router)
}
