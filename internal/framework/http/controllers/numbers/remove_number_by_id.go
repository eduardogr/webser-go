package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
)

func NewRemoveNumberByIdController(usecase usecases.RemoveNumberByIdUsecase) *RemoveNumberByIdController {
	return &RemoveNumberByIdController{
		RemoveNumberByIdUsecase: usecase,
	}
}

type RemoveNumberByIdController struct {
	RemoveNumberByIdUsecase usecases.RemoveNumberByIdUsecase
}

func (c *RemoveNumberByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	err := c.RemoveNumberByIdUsecase.Execute(id)

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

	response := map[string]interface{}{
		"success": true,
		"data": map[string]string{
			"message": "Number removed correctly",
		},
	}
	json.NewEncoder(w).Encode(response)
}
