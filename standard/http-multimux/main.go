package main

import (
	"log"
	"net/http"
)

func main() {
	submux := http.NewServeMux()
	submux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	submux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create user"))
	})

	mainmux := http.NewServeMux()
	mainmux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All users"))
	})

	mainmux.Handle("/", submux)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mainmux)
}
