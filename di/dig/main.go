package main

import (
	"log"
)

func main() {
	// Build the dependency injection container
	container := BuildContainer()

	// Register the App constructor
	if err := container.Provide(NewApp); err != nil {
		log.Fatalf("Failed to provide App: %v", err)
	}

	// Invoke the app and start the server
	err := container.Invoke(func(app *App) error {
		return app.Start()
	})
	if err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}