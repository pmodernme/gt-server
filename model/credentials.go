package model

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

func Signin(creds *Credentials) (string, error) {
	result := &Credentials{}
	if err := DB.Where("username = ?", creds.Username).First(result).Error; err != nil {
		return "", ErrUnknownUser
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(creds.Password))
	if err != nil {
		return "", ErrIncorrectPassword
	}

	expiration := time.Now().Add(time.Minute * 100000).Unix()
	tk := &Token{
		UserID: result.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", ErrInternal
	}

	return tokenString, nil
}
