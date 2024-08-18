package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All users"))
	})
	http.HandleFunc("GET /users/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User name: "))
		w.Write([]byte(r.PathValue("name")))
	})
	http.HandleFunc("GET /wild/{path...}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("path")))
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
