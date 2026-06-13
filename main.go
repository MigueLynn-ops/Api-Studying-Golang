package main

import (
	"Api-Aula_1/config"
	"Api-Aula_1/router"
	"log"
	"net/http"
	"os"
)

func main() {

	config.LoadEnv()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY not defined. Create a .env file based on the .env-example file.")
	}

	r := router.New()
	log.Printf("Starting server on %s\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r))
}
