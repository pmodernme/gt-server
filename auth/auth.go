package auth

import (
	"context"
	"encoding/json"
	"log"
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

	tk := r.Context().Value(Key("user")).(*model.Token)
	creds, err := model.GetCredsFromID(tk.UserID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(creds.Username)
}
