package main

import (
	"log"
	"net/http"
)

const PORT = "4000"

func main() {
	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Println("Listening on port", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(showSnippetRoute, showSnippet)
	mux.HandleFunc(createSnippetRoute, createSnippet)
	mux.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}
