package main

import (
	"fmt"

	"github.com/samber/do/v2"
)

func BuildContainer() do.Injector {
	injector := do.New()

	// Register providers using do.Provide with correct signatures
	do.Provide(injector, func(i do.Injector) (DSN, error) {
		return ProvideDSN(), nil
	})
	do.Provide(injector, func(i do.Injector) (Database, error) {
		dsn := do.MustInvoke[DSN](i)
		return NewDatabase(string(dsn)), nil
	})
	do.Provide(injector, func(i do.Injector) (UserService, error) {
		db := do.MustInvoke[Database](i)
		return NewUserService(db), nil
	})
	do.Provide(injector, func(i do.Injector) (*Server, error) {
		userService := do.MustInvoke[UserService](i)
		return NewServer(userService), nil
	})
	do.Provide(injector, func(i do.Injector) (string, error) {
		return ProvidePort(), nil
	})

	fmt.Println("samber/do container built successfully!")
	return injector
}

type App struct {
	Server *Server
	Port   string
}

func NewApp(injector do.Injector) *App {
	fmt.Println("Creating App with Server and Port dependencies")
	server := do.MustInvoke[*Server](injector)
	port := do.MustInvoke[string](injector)
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