package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!")) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All users")) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	http.HandleFunc("GET /users/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User name: "))       //nolint:errcheck,gosec // HTTP response write errors aren't useful
		w.Write([]byte(r.PathValue("name"))) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	http.HandleFunc("GET /wild/{path...}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("path"))) //nolint:errcheck,gosec // HTTP response write errors aren't useful
	})
	fmt.Println("Server is running on port 8080")
	sv := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(sv.ListenAndServe())
}
