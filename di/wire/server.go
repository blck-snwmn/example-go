package main

import (
	"fmt"
	"net/http"
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
		
		fmt.Fprintf(w, "Users: %v\n", users)
	})
	
	return http.ListenAndServe(":"+string(s.port), nil)
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