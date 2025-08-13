package main

import (
	"log"
)

func main() {
	// Build the dependency injection container
	injector := BuildContainer()

	// Create and start the app
	app := NewApp(injector)
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}