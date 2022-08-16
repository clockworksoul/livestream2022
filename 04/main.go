package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, GET!")
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, PUT!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the vars from the request
	vars := mux.Vars(r)

	// Get the one that we care about
	word := vars["word"]

	fmt.Fprintln(w, word)
}

func main() {
	// Create a new mux router
	r := mux.NewRouter()

	// Associate a path with a handler function on the router
	r.HandleFunc("/livestream/{word}",
		echoHandler).Methods("GET")

	r.HandleFunc("/method/", getHandler).Methods("GET")
	r.HandleFunc("/method/", putHandler).Methods("PUT")

	// Bind to a port and pass in the mux router
	log.Fatal(http.ListenAndServe(":8080", r))
}
