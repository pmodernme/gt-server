package API

import (
	"net/http"
)

// Mux for the API
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/users", users)
	return m
}
