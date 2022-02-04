package main

import (
	"flag"
	"net/http"

	"github.com/rhodeon/sniphub/pkg/prettylog"
)

const PORT = "4000"

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	setupHandlers(mux)

	prettylog.InfoF("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	prettylog.Error(err.Error())
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(showSnippetRoute, showSnippet)
	mux.HandleFunc(createSnippetRoute, createSnippet)
	mux.HandleFunc(staticRoute, serveStaticFiles)
}
