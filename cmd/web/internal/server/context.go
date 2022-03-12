package server

import (
	"github.com/rhodeon/sniphub/pkg/models"
	"net/http"
)

// ContextKey is used as an alias for request context keys.
type ContextKey string

const ContextKeyIsAuthenticated = ContextKey("isAuthenticated")
const ContextKeyUser = ContextKey("user")

// isAuthenticated returns true if the current request is from authenticated user,
// otherwise it returns false.
func (app *Application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(ContextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}

// getUserFromContext returns the current user data stored in a request's context.
// Returns an empty user struct on an error,
func (app *Application) getUserFromContext(r *http.Request) models.User {
	user, ok := r.Context().Value(ContextKeyUser).(models.User)
	if !ok {
		return models.User{}
	}
	return user
}
