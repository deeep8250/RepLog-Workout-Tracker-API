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
