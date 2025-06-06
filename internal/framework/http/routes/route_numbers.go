package routes

import controllers "github.com/eduardogr/webser-go/internal/framework/http/controllers/numbers"

type NumbersRoute struct {
	Path                       string
	PathID                     string
	CreateNumberController     *controllers.CreateNumberController
	GetAllNumbersController    *controllers.GetAllNumbersController
	GetNumberByIdController    *controllers.GetNumberByIdController
	UpdateNumberByIdController *controllers.UpdateNumberByIdController
	RemoveNumberByIdController *controllers.RemoveNumberByIdController
}

func NewNumbersRoute(
	createNumber *controllers.CreateNumberController,
	getAllNumbers *controllers.GetAllNumbersController,
	getNumberById *controllers.GetNumberByIdController,
	updateNumberById *controllers.UpdateNumberByIdController,
	removeNumberById *controllers.RemoveNumberByIdController,
) *NumbersRoute {
	return &NumbersRoute{
		Path:                       "/numbers",
		PathID:                     "/numbers/{id}",
		CreateNumberController:     createNumber,
		GetAllNumbersController:    getAllNumbers,
		GetNumberByIdController:    getNumberById,
		UpdateNumberByIdController: updateNumberById,
		RemoveNumberByIdController: removeNumberById,
	}
}
