package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewGetAllNumbersController(usecase usecases.GetAllNumbersUsecase) *GetAllNumbersController {
	return &GetAllNumbersController{
		GetAllNumbersUsecase: usecase,
	}
}

type GetAllNumbersController struct {
	GetAllNumbersUsecase usecases.GetAllNumbersUsecase
}

func (c *GetAllNumbersController) Execute(w http.ResponseWriter, r *http.Request) {
	numbers, err := c.GetAllNumbersUsecase.Execute()

	if err != nil {
		response := map[string]interface{}{
			"success": false,
			"data": map[string]string{
				"error": err.Error(),
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	numbersResponse := []domain.Number{}
	if numbers != nil {
		numbersResponse = numbers
	}

	response := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"numbers": numbersResponse,
		},
	}
	json.NewEncoder(w).Encode(response)
}
