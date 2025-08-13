//go:build wireinject

package main

import "github.com/google/wire"

func InitializeServer() (*Server, error) {
	wire.Build(
		ProvideDSN,
		NewDatabase,
		NewUserService,
		ProvidePort,
		NewServer,
	)
	return &Server{}, nil
}