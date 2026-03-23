package models

import "time"

type User struct {
	Id            int       `db:"id"`
	Email         string    `db:"email"`
	Password_hash string    `db:"password_hash"`
	Created_at    time.Time `db:"created_at"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type Exercises struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	MuscleGroup string    `db:"muscle_group" json:"muscle_group" binding:"requried"`
	CreatedBy   *int      `db:"created_by" json:"created_by" binding:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" binding:"required"`
}
