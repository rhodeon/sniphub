package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT = "4000"

// default home response
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to Texthub."))
}

// shows an example snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippet %d", id)
}

// allows user to create a snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

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
