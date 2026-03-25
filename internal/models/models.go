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
	Name        string    `db:"name" json:"name" binding:"required"`
	MuscleGroup string    `db:"muscle_group" json:"muscle_group" binding:"required"`
	CreatedBy   *int      `db:"created_by" json:"created_by"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" `
}

type Workouts struct {
	ID     int       `db:"id" json:"id"`
	UserID int       `db:"user_id" json:"user_id"`
	Notes  string    `db:"notes" json:"notes" binding:"required"`
	Date   time.Time `db:"date" json:"date"`
}

type Sets struct {
	ID         int    `db:"id" json:"id"`
	WorkoutID  int    `db:"workout_id" json:"workout_id"`
	ExerciseID int    `db:"exercise_id" json:"exercise_id" binding:"required"`
	Reps       int    `db:"reps" json:"reps" binding:"required"`
	Duration   int    `db:"duration" json:"duration"`
	Weight_kg  string `db:"weight_kg" json:"weight_kg" binding:"required"`
}
