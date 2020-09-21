package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitialRouter() *mux.Router {
	// Main Router
	router := mux.NewRouter()
	// Home Api register
	router.HandleFunc("/", MainApi)
	// 404 handler
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	// If you want to serve
	// CSS, JS & Images Statically.
	// uncomment this <
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	// return Mux router
	return router
}
