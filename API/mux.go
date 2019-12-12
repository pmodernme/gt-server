package API

import (
	"log"
	"net/http"
	"os"
)

// Port for the api to serve from
var Port string

func init() {
	Port = os.Getenv("API_PORT")
	if Port == "" {
		Port = "4040"
		log.Printf("API_PORT defaulting to %s", Port)
	}
}

// Mux for the API
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/users", users)
	return m
}
