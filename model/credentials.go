package model

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

var ErrUnknownUser = errors.New("model: no user exists with that username")
var ErrIncorrectPassword = errors.New("model: incorrect password for user")
var ErrInternal = errors.New("model: internal error")

func Signup(creds *Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	_, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))
	if err != nil {
		return ErrInternal
	}

	return nil
}

func Signin(creds *Credentials) error {
	result := db.QueryRow("select password from users where username=$1", creds.Username)

	storedCreds := &Credentials{}
	if err := result.Scan(&storedCreds.Password); err != nil {
		if err == sql.ErrNoRows {
			return ErrUnknownUser
		}

		return ErrInternal
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password))
	if err != nil {
		return ErrIncorrectPassword
	}

	return nil
}
