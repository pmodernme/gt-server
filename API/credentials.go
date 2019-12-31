package API

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
)

// Signup - API endpoint for signing up a new user
// Status 200 indicates success
func Signup(w http.ResponseWriter, r *http.Request) {
	creds := decodeCredentials(w, r)
	if err := model.Signup(creds); err != nil {
		var resp = map[string]interface{}{"status": false, "message": err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
	}
}

// Signin - API endpoint for signing in a user
<<<<<<< HEAD
// Success writes a JSON body including the username and a token
=======
// Success write a JSON body including the username and a token
>>>>>>> 9ac2288bd9efeaffa19338ba39c6363c7cda8d29
func Signin(w http.ResponseWriter, r *http.Request) {
	creds := decodeCredentials(w, r)

	userID, err := model.Signin(creds)
	if err != nil {
		var code int
		switch err {
		case model.ErrUnknownUser:
			code = http.StatusUnauthorized
		case model.ErrIncorrectPassword:
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}
		writeError(w, code, "Invalid Login Credentials. Please try again")
		return
	}

	expiration := time.Now().Add(time.Minute * 100000).Unix()
	tk := &model.Token{
		UserID: userID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error creating Token")
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  false,
		"message": "logged in",
		"token":   tokenString,
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

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  false,
		"message": message,
	})
}

// TestAuth -
func TestAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("It worked!")
}
