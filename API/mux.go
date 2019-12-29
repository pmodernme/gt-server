package API

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

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
	m.HandleFunc("/signup", Signup)
	return m
}
