package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"performance/models"
	"performance/utils"
)

func MainApi(w http.ResponseWriter, r *http.Request) {
	response := json.NewEncoder(w).Encode(utils.SuccessResponse(models.Data{
		Name:    "Munis",
		Surname: "Isazade",
	}))
	if response != nil {
		fmt.Println("Error response")
	}
}

