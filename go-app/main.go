package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		fmt.Fprintf(w, "Welcome to root web")
	} else {
		fmt.Fprintf(w, "you access the %s", path)
	}
}

func main() {
	http.HandleFunc("/", handler)

	// Create a server with timeouts
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // use default handler
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}