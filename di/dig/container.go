package main

import (
	"fmt"
	"log"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Register providers in dependency order
	if err := container.Provide(ProvideDSN); err != nil {
		log.Fatalf("Failed to provide DSN: %v", err)
	}

	if err := container.Provide(func(dsn DSN) Database {
		return NewDatabase(string(dsn))
	}); err != nil {
		log.Fatalf("Failed to provide Database: %v", err)
	}

	if err := container.Provide(NewUserService); err != nil {
		log.Fatalf("Failed to provide UserService: %v", err)
	}

	if err := container.Provide(NewServer); err != nil {
		log.Fatalf("Failed to provide Server: %v", err)
	}

	if err := container.Provide(ProvidePort); err != nil {
		log.Fatalf("Failed to provide Port: %v", err)
	}

	fmt.Println("Dig container built successfully!")
	return container
}

type App struct {
	Server *Server
	Port   string
}

func NewApp(server *Server, port string) *App {
	fmt.Println("Creating App with Server and Port dependencies")
	return &App{
		Server: server,
		Port:   port,
	}
}

func (a *App) Start() error {
	fmt.Println("Starting application...")
	a.Server.Routes()
	return a.Server.Start(a.Port)
}
