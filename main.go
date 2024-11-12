package main

import (
	"build_tag_practice/internal/logger"
	"log/slog"
	"net/http"
)

func main() {
	// ロガーの初期化
	logger.Setup()

	http.HandleFunc("/", handler)
	slog.Info("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "received request",
		"path", r.URL.Path,
		"method", r.Method,
	)
	w.Write([]byte("Hello, World!"))
}
