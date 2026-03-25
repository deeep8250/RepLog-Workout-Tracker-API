package services

import "replog/internal/models"

type SetsServiceInterface interface {
	CreateSets(sets *models.Sets) error
}
