package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewGetNumberByIdController(usecase usecases.GetNumberByIdUsecase) *GetNumberByIdController {
	return &GetNumberByIdController{
		GetNumberByIdUsecase: usecase,
	}
}

type GetNumberByIdController struct {
	GetNumberByIdUsecase usecases.GetNumberByIdUsecase
}

func (c *GetNumberByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	number, err := c.GetNumberByIdUsecase.Execute(id)

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

	if number != (domain.Number{}) {
		response := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"number": number,
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}
