package main

import (
	"log"
	"net/http"

	"./API"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file.", err)
	}
}

func main() {
	http.Handle("/", API.Handlers())
	log.Printf("API listening on port %s", API.Port)
	log.Fatal(http.ListenAndServe(":"+API.Port, nil))
}
