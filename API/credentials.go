package API

import (
	"net/http"
	"time"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
)

// Signup - API endpoint for signing up a new user
// Status 200 indicates success
func Signup(w http.ResponseWriter, r *http.Request) {
	var creds *model.Credentials
	decode(creds, w, r)
	if err := model.Signup(creds); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

// Signin - API endpoint for signing in a user
// Success writes a JSON body including the username and a token
func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &model.Credentials{}
	decode(creds, w, r)

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

	send(map[string]interface{}{
		"token": tokenString,
		"user":  creds.Username,
	}, true, "logged in successfully", w)
}
