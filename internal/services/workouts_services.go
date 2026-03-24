package services

import (
	"replog/internal/models"
	"replog/internal/repository"
)

type WorkoutsServices struct {
	repoWorkouts *repository.WorkoutsRepo
}

func NewWorkoutsService(WorkoutRepo *repository.WorkoutsRepo) *WorkoutsServices {
	return &WorkoutsServices{repoWorkouts: WorkoutRepo}
}

func (s *WorkoutsServices) CreateWorkouts(workouts *models.Workouts) error {
	err := s.repoWorkouts.CreateWorkouts(workouts)
	if err != nil {
		return err
	}
	return nil
}

func (s *WorkoutsServices) GetWorkouts(userID int) ([]models.Workouts, error) {

	workouts, err := s.repoWorkouts.GetWorkouts(userID)
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

func (s *WorkoutsServices) GetWorkoutsByID(ID, userID int) (*models.Workouts, error) {

	workouts, err := s.repoWorkouts.GetWorkoutsByID(ID, userID)
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

func (s *WorkoutsServices) DeleteWorkouts(ID, userID int) error {

	err := s.repoWorkouts.DeleteWorkoutByID(ID, userID)
	if err != nil {
		return err
	}
	return nil
}
