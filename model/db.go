package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Users by ID
var Users map[string]User

func init() {
	LoadUsers()
}

// AllUsers as a slice
func AllUsers() []User {
	us := make([]User, 0)
	for _, u := range Users {
		us = append(us, u)
	}
	return us
}

// GetUsers in a slice of ids
func GetUsers(ids []string) map[string]User {
	us := make(map[string]User, 0)
	for _, id := range ids {
		us[id] = Users[id]
	}
	return us
}

// SaveUsers to disc
func SaveUsers() {
	storeJSON(Users, "user_data")
}

// LoadUsers from disc or create new user DB
func LoadUsers() {
	err := loadJSON(&Users, "user_data")
	if err != nil {
		fmt.Println("Creating new User DB")
		Users = make(map[string]User)
	}
}

func storeJSON(j interface{}, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(j)
}

func loadJSON(j interface{}, fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&j)
	if err != nil {
		return err
	}

	return nil
}
