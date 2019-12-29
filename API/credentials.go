package API

import (
	"encoding/json"
	"net/http"

	"../model"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = model.Signup(creds); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// func Signin
