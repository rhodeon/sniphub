package main

import (
	"log"
	"net/http"
)

const PORT = "4000"

// default home response
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Texthub"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Listening on port", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
