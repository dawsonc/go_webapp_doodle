package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	// Set up handlers
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Find the port we're supposed to bind to
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080
	}

	// Engage the webserver
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
