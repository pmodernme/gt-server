package API

import (
	"encoding/json"
	"net/http"

	"../model"
)

// UserRequest for user data by ID
// IDs a slice of user ID as string
type UserRequest struct {
	IDs []string `json:"ids"`
}

func users(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getUsers(w, r)
		return
	} else if r.Method == http.MethodPost {
		putUser(w, r)
		return
	} else {
		http.Error(w, http.StatusText(400), 400)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	jq := r.URL.Query().Get("json")
	if jq == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(model.AllUsers())
		return
	}

	rd := UserRequest{}
	err := json.Unmarshal([]byte(jq), &rd)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GetUsers(rd.IDs))
}

func putUser(w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	model.Users[u.ID] = u
	model.SaveUsers()
}
