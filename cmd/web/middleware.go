package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"

	"github.com/rhodeon/sniphub/pkg/prettylog"
)

// requireAuthentication redirects a user to the login page if they aren't logged in.
func (app *application) requireAuthentication(next http.Handler) http.Handler {
	fmt.Println("authhhhh")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// if the user is not authenticated, redirect them to the login page and
			// return from the middleware chain so that no subsequent handlers in
			// the chain are executed.
			if !app.isAuthenticated(r) {
				http.Redirect(w, r, "/user/login", http.StatusSeeOther)
				return
			}

			// otherwise, set the "Cache-Control: no-store" header so that pages
			// require authentication are not stored in the user's browser cache (or
			// other intermediary cache).
			w.Header().Add("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
		},
	)
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("X-Frame-Options", "deny")
			rw.Header().Set("X-XSS-Protection", "1; mode=block")
			next.ServeHTTP(rw, r)
		},
	)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			prettylog.InfoF("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
			next.ServeHTTP(rw, r)
		},
	)
}

func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					rw.Header().Set("Connection", "close")
					serverError(rw, fmt.Errorf("%s", err))
				}
			}()

			next.ServeHTTP(rw, r)
		},
	)
}

// noSurf is a middleware which uses a customized CSRF cookie with
// the Secure, Path and HttpOnly flags set.
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
