package handler

import (
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	// We only accept POST requests because the client is sending data to create a new short URL
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("only POST allowed"))
		return
	}

	// Parse the incoming form data so we can read the fields sent in the request body
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid form"))
		return
	}

	// Extract the "url" field from the submitted form data
	longURL := r.FormValue("url")

	// Check if the user actually provided a URL to shorten, stop if the field is empty
	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing url"))
		return
	}

	// Placeholder response to confirm we successfully extracted the URL
	w.Write([]byte("received: " + longURL))
}
