package main

import (
	"github.com/google/wire"

	"github.com/eduardogr/webser-go/pkg/api"
	"github.com/eduardogr/webser-go/pkg/repository"
	"github.com/eduardogr/webser-go/pkg/server"
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
