package main

import (
	"net/http"

	"github.com/rhodeon/sniphub/pkg/prettylog"
)

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
