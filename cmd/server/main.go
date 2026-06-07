package main

import (
	"log/slog"
	"os"
)

func main() {
	// Set up a JSON logger to print to the terminal so we get clean, structured logs
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(handler))

	// quick log statement to confirm if logging is working.
	slog.Info("starting url-shortener server")
}
