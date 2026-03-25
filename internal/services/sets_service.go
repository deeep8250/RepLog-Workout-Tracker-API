package services

import (
	"replog/internal/models"
	"replog/internal/repository"
)

type SetsService struct {
	repo *repository.SetsRepo
}

func NewSetService(Repo *repository.SetsRepo) *SetsService {
	return &SetsService{
		repo: Repo,
	}
}

func (s *SetsService) CreateSets(sets *models.Sets) error {
	err := s.repo.CreateSets(sets)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetsService) GetAllSets(workoutID int) ([]models.Sets, error) {
	results, err := s.repo.GetSets(workoutID)
	if err != nil {
		return nil, err
	}
	return results, nil
}
