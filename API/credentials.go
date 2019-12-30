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
		var resp = map[string]interface{}{"status": false, "message": err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := decodeCredentials(w, r)

	token, err := model.Signin(creds)
	if err != nil {
		switch err {
		case model.ErrUnknownUser:
			w.WriteHeader(http.StatusUnauthorized)
		case model.ErrIncorrectPassword:
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Invalid Login Credentials. Please try again",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  false,
		"message": "logged in",
		"token":   token,
		"user":    creds.Username,
	})
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

func TestAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("It worked!")
}
