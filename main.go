package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"performance/config"
	"performance/routes"
	"performance/utils"
)

func main() {
	logger := utils.Logger{}
	logger.Fatal("Error")
	// first load application environment variables
	e := godotenv.Load() //Load default .env file
	if e != nil {
		// if problem problem
		log.Panic(e)
	}
	r := routes.InitialRouter()
	fmt.Printf("Starting application at: http://localhost%s\n", config.Port)
	lis := http.ListenAndServe(config.Port, utils.MiddlewareRequest(r))
	if lis != nil {
		println("Error starting app")
	}

}
