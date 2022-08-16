package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello gorilla/mux!")
}

func main() {
	// Associate a path with a handler function
	http.HandleFunc("/", helloGoHandler)

	// Bind to a port and default to the DefaultServeMux handler.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
