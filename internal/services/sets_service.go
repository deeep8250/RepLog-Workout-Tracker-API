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

func Calculate1RMService(weight, reps float64) float64 {

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

func (s *SetsService) ProgressReportService(userID, exerciseID int) ([]models.ProgressEntry, error) {
	value, err := s.repo.ProgressReport(userID, exerciseID)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *SetsService) ShouldIncreaseService(targetReps, userID int) ([]models.ShouldIncreseOverload, error) {

	var returnResults []models.ShouldIncreseOverload
	var should bool

	// getting all exercises created by the user
	exercises, err := s.repo.ShouldIncreaseRepsRepo(userID)
	if err != nil {
		return nil, err
	}

	// it first get the last 3 sets accroding to the exercises and then
	// pass those exercise as list of slice in the shouldincreaseweight function and then it return back the overload result
	for _, exercise := range exercises {

		v, err := s.repo.GetLastThreeSets(userID, exercise.ID)
		if err != nil {
			return nil, err
		}
		should = ShouldIncreaseWeight(v, targetReps)

		report := models.ShouldIncreseOverload{
			ExerciseName:  exercise.Name,
			ExerciseID:    exercise.ID,
			ShouldIncrese: should,
		}
		returnResults = append(returnResults, report)

	}

	return returnResults, nil

}
