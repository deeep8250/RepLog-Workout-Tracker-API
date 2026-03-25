package services

import "replog/internal/models"

type WorkoutServiceInterface interface {
	GetWorkoutsByID(workoutID, userID int) (*models.Workouts, error)
}
