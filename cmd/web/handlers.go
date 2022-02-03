package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// default home response
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to Sniphub."))
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
