package API

import (
	"encoding/json"
	"net/http"
)

func decode(obj interface{}, w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Error decoding json object")
	}
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    false,
		"message":    message,
		"error_code": code,
		"data":       map[string]interface{}{},
	})
}

func send(data map[string]interface{}, success bool, message string, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data,
	})
}
