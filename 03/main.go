package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello gorilla/mux!")
}

func main() {
	// Create a new mux router
	r := mux.NewRouter()

	// Associate a path with a handler function on the router
	r.HandleFunc("/", helloMuxHandler)

	// Bind to a port and pass in the mux router
	log.Fatal(http.ListenAndServe(":8080", r))
}
