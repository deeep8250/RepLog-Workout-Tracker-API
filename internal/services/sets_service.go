package services

import (
	"fmt"
	"replog/internal/models"
	"replog/internal/repository"
	"strconv"
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

func (s *SetsService) DeleteFromSets(workoutID, setsID int) error {
	err := s.repo.DeleteSets(workoutID, setsID)
	if err != nil {
		return err
	}
	return nil
}

func Calculate1RM(weight, reps float64) float64 {

	rm := weight * (1 + reps/30)
	return rm

}

func (s *SetsService) TotalVolume(userID, exerciseID int) (float64, error) {

	lastThreeSets, err := s.repo.GetLastThreeSets(userID, exerciseID)
	if err != nil {
		return 0, err
	}

	var totalVolumes float64
	for _, value := range lastThreeSets {
		WeightInt, err := strconv.ParseFloat(value.Weight_kg, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue

		}

		tempV := WeightInt * float64(value.Reps)
		totalVolumes += float64(tempV)

	}

	return totalVolumes, nil
}
