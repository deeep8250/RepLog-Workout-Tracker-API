package services

import (
	"fmt"
	"replog/internal/models"
	"replog/internal/repository"
)

type ExerciseServices struct {
	repoExerCise *repository.ExerciseRepo
}

func NewExerciseService(ExerCiseRepo *repository.ExerciseRepo) *ExerciseServices {
	return &ExerciseServices{repoExerCise: ExerCiseRepo}
}

func (s *ExerciseServices) GetAllExercises(muscle string, id int) ([]models.Exercises, error) {

	exercises, err := s.repoExerCise.GetAllExercises(muscle, id)
	if err != nil {
		return nil, err
	}

	fmt.Println("exercises : ", exercises)

	return exercises, nil
}

func (s *ExerciseServices) CreateExercises(exercises *models.Exercises) error {

	err := s.repoExerCise.InsertExercise(exercises)
	if err != nil {
		return err
	}
	return nil

}

// func (s *ExerciseServices) CreateWorkouts(workouts *models.Workouts) error {
// 	err := s.repoExerCise.CreateWorkouts(workouts)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *ExerciseServices) GetWorkouts(userID int) ([]models.Workouts, error) {

// 	workouts, err := s.repoExerCise.GetWorkouts(userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return workouts, nil
// }
