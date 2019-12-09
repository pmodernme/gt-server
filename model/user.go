package model

// User for public consumption
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
