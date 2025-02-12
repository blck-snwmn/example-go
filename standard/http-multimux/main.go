package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	submux := http.NewServeMux()
	submux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!")) //nolint: errcheck
	})
	submux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create user")) //nolint: errcheck
	})

	mainmux := http.NewServeMux()
	mainmux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All users")) //nolint: errcheck
	})

	mainmux.Handle("/", submux)

	log.Println("Server is running on port 8080")
	sv := &http.Server{
		Addr:              ":8080",
		Handler:           mainmux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(sv.ListenAndServe())
}
