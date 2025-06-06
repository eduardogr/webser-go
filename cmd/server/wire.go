//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	usecases "github.com/eduardogr/webser-go/internal/application/usecases/numbers"
	"github.com/eduardogr/webser-go/internal/framework/http"
	controllers "github.com/eduardogr/webser-go/internal/framework/http/controllers/numbers"
	"github.com/eduardogr/webser-go/internal/framework/http/routes"
	"github.com/eduardogr/webser-go/internal/framework/storage/mysql"
)

var NumberRouteControllersSet = wire.NewSet(
	controllers.NewCreateNumberController,
	controllers.NewGetAllNumbersController,
	controllers.NewGetNumberByIdController,
	controllers.NewUpdateNumberByIdController,
	controllers.NewRemoveNumberByIdController,
)

var UseCasesSet = wire.NewSet(
	usecases.NewCreateNumberUsecase,
	usecases.NewGetAllNumbersUsecase,
	usecases.NewGetNumberByIdUsecase,
	usecases.NewUpdateNumberByIdUsecase,
	usecases.NewRemoveNumberByIdUsecase,
)

func InitializeServer() (*http.Server, error) {
	wire.Build(
		mysql.SetupMysqlDatabase,
		mysql.NewMysqlNumberRepository,
		UseCasesSet,
		NumberRouteControllersSet,
		routes.NewStatusRoute,
		routes.NewNumbersRoute,
		routes.NewRouterBuilder,
		http.NewServer,
	)
	return &http.Server{}, nil
}
