package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/rhodeon/prettylog"
	"github.com/rhodeon/sniphub/cmd/web/internal/session"
	"github.com/rhodeon/sniphub/pkg/models"
	"net/http"
)

// requireAuthentication redirects a user to the login page if they aren't logged in.
func (app *Application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// if the user is not authenticated, redirect them to the login page and
			// return from the middleware chain so that no subsequent handlers in
			// the chain are executed.
			if !app.isAuthenticated(r) {
				app.SessionManager.Put(r.Context(), session.KeyFlashMessage, session.LoginPrompt)
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
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

// secureHeaders sets generic protection headers
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

// authenticate sets the request context with a valid authentication if it comes from a valid user.
// Otherwise, it passes on the request as-is.
func (app Application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// carry on without authenticating if there is currently no user id key set
			exists := app.SessionManager.Exists(r.Context(), session.KeyUserId)
			if !exists {
				next.ServeHTTP(w, r)
				return
			}

			// attempt to fetch user details if the user id key exists
			user, err := app.Users.Get(app.SessionManager.GetInt(r.Context(), session.KeyUserId))
			if err != nil {
				if errors.Is(err, models.ErrInvalidUser) {
					// remove user id key and proceed if no user is found
					app.SessionManager.Remove(r.Context(), session.KeyUserId)
					next.ServeHTTP(w, r)
				} else {
					// raise an error for any other reason
					serverError(w, err)
				}
				return
			}

			// do not authenticate if the user is inactive
			if !user.Active {
				app.SessionManager.Remove(r.Context(), session.KeyUserId)
				next.ServeHTTP(w, r)
				return
			}

			// otherwise, authenticate the current request for future handlers
			ctx := context.WithValue(r.Context(), ContextKeyIsAuthenticated, true)
			ctx = context.WithValue(ctx, ContextKeyUser, user) // save user data in context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
