package main

import (
	"flag"
	"net/http"

	"github.com/rhodeon/sniphub/pkg/logging"
)

const PORT = "4000"

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	setupHandlers(mux)

	logging.InfoF("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	logging.Error(err.Error())
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(showSnippetRoute, showSnippet)
	mux.HandleFunc(createSnippetRoute, createSnippet)
	mux.HandleFunc(staticRoute, serveStaticFiles)
}
