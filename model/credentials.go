package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Credentials - stored in SQL
type Credentials struct {
	gorm.Model
	Password string `json:"password"` //bcrypt hash, not actual password
	Username string `gorm:"unique_index" json:"username"`
}

// ErrUnknownUser - The error returned if the username is incorrect
var ErrUnknownUser = errors.New("model: no user exists with that username")

// ErrIncorrectPassword - The error return if the password is incorrect
var ErrIncorrectPassword = errors.New("model: incorrect password for user")

// ErrInternal - The error returned if the fault of the error is in code
var ErrInternal = errors.New("model: internal error")

// Signup - returns nil if successful, an error if not
func Signup(creds *Credentials) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrInternal
	}

	creds.Password = string(pass)

	openDB()
	defer DB.Close()

	newCreds := DB.Create(creds)
	if newCreds.Error != nil {
		return newCreds.Error
	}

	return nil
}

// Signin - Returns the UserID or an Error
func Signin(creds *Credentials) (uint, error) {
	result := &Credentials{}

	openDB()
	defer DB.Close()

	if err := DB.Where("username = ?", creds.Username).First(result).Error; err != nil {
		return 0, ErrUnknownUser
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(creds.Password))
	if err != nil {
		return 0, ErrIncorrectPassword
	}

	return result.ID, nil
}
