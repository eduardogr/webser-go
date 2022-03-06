package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/eduardogr/webser-go/pkg/api"
)

type Handler struct {
	numberService api.NumberService
}

func NewHandler(n api.NumberService) *Handler {
	return &Handler{
		numberService: n,
	}
}

func (h *Handler) GetAllNumbers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: allNumbersHandler")

	Numbers, err := h.numberService.GetAll()

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

	numbersResponse := []api.Number{}
	if Numbers != nil {
		numbersResponse = Numbers
	}

	response := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"numbers": numbersResponse,
		},
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: singleNumberHandler")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	number, err := h.numberService.Get(id)

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

	if number != (api.Number{}) {
		response := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"number": number,
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}

func (h *Handler) CreateNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewNumber")

	// TODO: check that number does NOT exist

	// get the body of our POST request
	// unmarshal this into a new NewNumberRequest struct

	reqBody, _ := ioutil.ReadAll(r.Body)

	var n api.NewNumberRequest
	err := json.Unmarshal(reqBody, &n)
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

	// create request
	err = h.numberService.New(n)

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
			"number": n.ID,
		},
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateNumberHandler")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var n api.Number
	json.Unmarshal(reqBody, &n)

	err := h.numberService.Update(n, id)

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

func (h *Handler) DeleteNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteNumber")

	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	err := h.numberService.Remove(id)

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
