package main

import (
	"log"
	"net/http"
)

const PORT = "4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", showSnippet)
	mux.HandleFunc("/snippet/create/", createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Listening on port", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
