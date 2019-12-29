package main

import (
	"fmt"
	"log"
	"net/http"

	"./API"
	"github.com/joho/godotenv"

	"./model"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file.", err)
	}

	model.InitDB()
}

func main() {
	log.Printf("API listening on port %s", API.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", API.Port), API.Mux())
}
