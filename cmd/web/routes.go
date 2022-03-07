package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	homeRoute        = "/"
	showSnipRoute    = "/snip"
	createSnipRoute  = "/snip/create"
	staticRoute      = "/static/*"
	latestSnipsRoute = "/latest"
)

func (app *application) routesHandler() http.Handler {
	router := chi.NewRouter()

	// set middleware
	router.Use(recoverPanic, logRequests, secureHeaders)
	router.Use(app.sessionManager.LoadAndSave)
	router.Use(middleware.CleanPath, middleware.StripSlashes)

	// set route handlers
	router.Get("/", app.home)
	router.Get("/static/*", app.serveStaticFiles)
	router.Get("/latest", app.showLatestSnips)

	router.Route("/snip", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", app.showSnip)
		r.Get("/create", app.createSnipGet)
		r.Post("/create", app.createSnipPost)
	})

	router.Route("/user", func(r chi.Router) {
		r.Get("/signup", app.signupUserGet)
		r.Post("/signup", app.signupUserPost)
		r.Get("/login", app.loginUserGet)
		r.Post("/login", app.loginUserPost)

		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthentication)
			r.Post("/logout", app.logoutUser)
		})
	})

	return router
}
