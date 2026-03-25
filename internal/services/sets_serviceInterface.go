package services

import "replog/internal/models"

type SetsServiceInterface interface {
	CreateSets(sets *models.Sets) error
	GetAllSets(workoutID int) ([]models.Sets, error)
}
