package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registering a Request Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested this %s path", r.URL.Path)
	})

	// Accept connection
	http.ListenAndServe(":80", nil)
}
