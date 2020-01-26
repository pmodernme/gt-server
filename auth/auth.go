package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"../model"

	jwt "github.com/dgrijalva/jwt-go"
)

type Exception model.Exception
type Key string

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &model.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), Key("user"), tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Test - returns response code 418 (Teapot) if good
func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

// GetUsername - returns the current username or an error if the token in invalid
func GetUsername(r *http.Request) (string, error) {
	if tk, ok := r.Context().Value(Key("user")).(*model.Token); ok && tk.Username != "" {
		return tk.Username, nil
	}

	return "", errors.New("Invalid Token")
}
