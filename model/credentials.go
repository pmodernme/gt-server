package model

import "golang.org/x/crypto/bcrypt"

type Credentials struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

func Signup(creds *Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	_, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))

	return err
}
