package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

	// set middleware
	router.Use(recoverPanic, logRequests, secureHeaders)

	// set route handlers
	router.Get(homeRoute, app.home)
	router.Get(showSnipRoute, app.showSnip)
	router.Post(createSnipRoute, app.createSnip)
	router.Get(staticRoute, app.serveStaticFiles)
	router.Get(latestSnipsRoute, app.showLatestSnips)

	return router
}
