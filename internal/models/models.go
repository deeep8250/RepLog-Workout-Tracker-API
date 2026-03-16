package models

import "time"

type User struct {
	Id            int       `db:"id"`
	Email         string    `db:"email"`
	Password_hash string    `db:"password_hash"`
	Created_at    time.Time `db:"created_at"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Passowrd string `json:"passowrd"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
