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

func (s *ExerciseServices) GetAllExercises(id int) ([]models.Exercises, error) {

	exercises, err := s.repoExerCise.GetAllExercises(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("exercises : ", exercises)

	return exercises, nil
}
