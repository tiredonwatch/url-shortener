package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// Set up a JSON logger to print to the terminal for logs
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(handler))

	mux := http.NewServeMux()

	// Call this endpoint to verify if the server is running
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Middleware intercepts incoming requests to log their duration
	interceptingMux := loggingMiddleware(mux)

	slog.Info("starting server", "port", 8080)

	// error handling and web server starting
	err := http.ListenAndServe(":8080", interceptingMux) // "All Interfaces:port 8080"
	if err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}

// This intercepts incoming requests so we can time them and log the details
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start a timer right when the request hits the server
		start := time.Now()

		// Let the request continue on to its handler function
		next.ServeHTTP(w, r)

		// Log the HTTP method, URL path, and how many milliseconds it took to finish
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
