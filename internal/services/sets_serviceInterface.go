package services

import "replog/internal/models"

type SetsServiceInterface interface {
	CreateSets(sets *models.Sets) error
	GetAllSets(workoutID int) ([]models.Sets, error)
	DeleteFromSets(workoutID, setsID int) error
	ProgressReportService(userID, exerciseID int) ([]models.ProgressEntry, error)
	ShouldIncreaseService(targetReps, userID int) ([]models.ShouldIncreseOverload, error)
}
