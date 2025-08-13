package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting Wire DI example...")
	
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	
	fmt.Println("Server initialized successfully with dependency injection!")
	fmt.Println("Visit http://localhost:8080/users to see the users")
	
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}