package main

import (
	"flag"
	"log"
	"net/http"
)

const PORT = "4000"

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Println("Starting server on", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(showSnippetRoute, showSnippet)
	mux.HandleFunc(createSnippetRoute, createSnippet)
	mux.HandleFunc(staticRoute, serveStaticFiles)
}
