package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewUpdateNumberByIdController(usecase usecases.UpdateNumberByIdUsecase) *UpdateNumberByIdController {
	return &UpdateNumberByIdController{
		UpdateNumberByIdUsecase: usecase,
	}
}

type UpdateNumberByIdController struct {
	UpdateNumberByIdUsecase usecases.UpdateNumberByIdUsecase
}

func (c *UpdateNumberByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := io.ReadAll(r.Body)
	var n domain.Number
	json.Unmarshal(reqBody, &n)

	err := c.UpdateNumberByIdUsecase.Execute(n, id)

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
		"data": map[string]interface{}{
			"number": n,
		},
	}
	json.NewEncoder(w).Encode(response)
}
