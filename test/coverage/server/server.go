package server

import "net/http"

func GenHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!")) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	mux.HandleFunc("GET /goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye, World!")) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	return mux
}
