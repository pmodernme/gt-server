package API

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := decodeCredentials(w, r)
	if err := model.Signup(creds); err != nil {
		log.Println("Internal Server Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := decodeCredentials(w, r)
	if err := model.Signin(creds); err != nil {
		log.Println("Internal Server Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func decodeCredentials(w http.ResponseWriter, r *http.Request) *model.Credentials {
	creds := &model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Println("Bad Request:", r)
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	return creds
}
