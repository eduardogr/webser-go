package main

import (
	"github.com/google/wire"

	"github.com/eduardogr/webser-go/internal/api"
	"github.com/eduardogr/webser-go/internal/repository"
	"github.com/eduardogr/webser-go/internal/server"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		repository.SetupDatabase,
		repository.NewStorage,
		repository.ProvideNumberRepository,
		api.NewNumberService,
		server.NewHandler,
		server.NewRouter,
		server.NewServer,
	)
	return &server.Server{}, nil
}
