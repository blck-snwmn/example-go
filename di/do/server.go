package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	userService UserService
}

func NewServer(userService UserService) *Server {
	fmt.Println("Creating Server with UserService dependency")
	return &Server{userService: userService}
}

func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetUsers(w, r)
	case http.MethodPost:
		s.handleCreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := s.userService.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string][]string{"users": users}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := s.userService.CreateUser(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) Start(port string) error {
	http.HandleFunc("/users", s.handleUsers)
	fmt.Printf("Server starting on port %s\n", port)
	
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	return server.ListenAndServe()
}

func (s *Server) Routes() {
	fmt.Println("Available routes:")
	fmt.Println("  GET  /users - List all users")
	fmt.Println("  POST /users - Create a new user (JSON body: {\"name\": \"username\"})")
}

func ProvidePort() string {
	return "8080"
}