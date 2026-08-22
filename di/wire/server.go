package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	userService UserService
	port        Port
}

func (s *Server) Start() error {
	fmt.Printf("Starting server on port %s\n", s.port)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := s.userService.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := fmt.Fprintf(w, "Users: %v\n", users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:              ":" + string(s.port),
		ReadHeaderTimeout: 5 * time.Second,
	}
	return server.ListenAndServe()
}

func NewServer(userService UserService, port Port) *Server {
	return &Server{
		userService: userService,
		port:        port,
	}
}

type Port string

func ProvidePort() Port {
	return Port("8080")
}
