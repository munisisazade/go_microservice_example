package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"performance/utils"
)

// 404 page not found handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	// set header status code to 404
	w.WriteHeader(http.StatusNotFound)
	// add response data
	tkn := json.NewEncoder(w).Encode(utils.NotFoundResponse())
	if tkn != nil {
		fmt.Println("Error ")
	}
}
