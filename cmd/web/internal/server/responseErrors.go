package server

import (
	"fmt"
	"github.com/rhodeon/prettylog"
	"net/http"
	"runtime/debug"
)

const ErrPageNotFound = "404 page not found"

// Prints stacktrace and sends internal server error to user.
func serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	prettylog.Error(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func notFoundError(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
