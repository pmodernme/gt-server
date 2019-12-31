package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	gorm.Model
	Password string `json:"password"`
	Username string `gorm:"unique_index" json:"username"`
}

var ErrUnknownUser = errors.New("model: no user exists with that username")
var ErrIncorrectPassword = errors.New("model: incorrect password for user")
var ErrInternal = errors.New("model: internal error")

func Signup(creds *Credentials) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrInternal
	}

	creds.Password = string(pass)

	newCreds := DB.Create(creds)
	if newCreds.Error != nil {
		return newCreds.Error
	}

	return nil
}

// Signin - Returns the UserID or an Error
func Signin(creds *Credentials) (uint, error) {
	result := &Credentials{}
	if err := DB.Where("username = ?", creds.Username).First(result).Error; err != nil {
		return 0, ErrUnknownUser
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(creds.Password))
	if err != nil {
		return 0, ErrIncorrectPassword
	}

	return result.ID, nil
}
