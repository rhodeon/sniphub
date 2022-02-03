package main

import (
	"log"
	"net/http"
)

const PORT = "4000"

// default home response
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Texthub."))
}

// shows an example snippet
func showSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("This is an example snippet."))
}

// allows user to create a snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can create a snippet here."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", showSnippet)
	mux.HandleFunc("/snippet/create/", createSnippet)

	log.Println("Listening on port", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
