package main

import (
	"fmt"
	"log"
	"net/http"

	"./API"
)

func main() {
	log.Printf("API listening on port %s", API.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", API.Port), API.Mux())
}
