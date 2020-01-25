package API

import (
	"log"
	"net/http"
	"os"

	"../auth"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // Autoload .env variables
)

// Port being used by the API
var Port string

func init() {
	Port = os.Getenv("API_PORT")
	if Port == "" {
		Port = "4040"
		log.Printf("API_PORT defaulting to %s", Port)
	}
}

// Routes - s handles routes that require authentication
func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/signin", Signin).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)

	s.HandleFunc("/test", auth.Test).Methods("GET")

	s.HandleFunc("/events", AllEvents).Methods("GET")
	s.HandleFunc("/events", NewEvent).Methods("POST")
	s.HandleFunc("/events", DeleteEvent).Methods("DELETE")
	s.HandleFunc("/events", UpdateEvent).Methods("PUT")

	return r
}

// CommonMiddleware - Adds JSON headers
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
