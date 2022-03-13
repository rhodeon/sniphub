package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	homeRoute   = "/"
	staticRoute = "/static/*"

	showSnipRoute   = "/snip"
	createSnipRoute = "/snip/create"
	userSnipsRoute  = "/user/{username}"

	signupRoute         = "/auth/signup"
	loginRoute          = "/auth/login"
	logoutRoute         = "/auth/logout"
	changePasswordRoute = "/auth/change-password"
)

func (app *Application) RouteHandler() http.Handler {
	router := chi.NewRouter()

	// set middleware
	router.Use(recoverPanic, logRequests, secureHeaders)
	router.Use(app.SessionManager.LoadAndSave)
	router.Use(middleware.CleanPath, middleware.StripSlashes)
	router.Use(noSurf, app.authenticate)

	// set route handlers
	router.Get(staticRoute, app.serveStaticFiles)
	router.Get(userSnipsRoute, app.showUserSnips)

	router.Route(homeRoute, func(r chi.Router) {
		r.Get("/", app.home)
		r.Get("/{page:[0-9]+}", app.home)
	})

	router.Route("/snip", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", app.showSnip)
		r.Get("/create", app.createSnipGet)
		r.Post("/create", app.createSnipPost)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Get("/signup", app.signupUserGet)
		r.Post("/signup", app.signupUserPost)
		r.Get("/login", app.loginUserGet)
		r.Post("/login", app.loginUserPost)
		r.Get("/change-password", app.changePasswordGet)
		r.Post("/change-password", app.changePasswordPost)

		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthentication)
			r.Post("/logout", app.logoutUser)
		})
	})

	return router
}
