package repository

import (
	"replog/internal/models"

	"github.com/jmoiron/sqlx"
)

type ExerciseRepo struct {
	db *sqlx.DB
}

func NewExerCiseRepo(Db *sqlx.DB) *ExerciseRepo {
	return &ExerciseRepo{db: Db}
}

func (r *ExerciseRepo) GetAllExercises(muscle string, id int) ([]models.Exercises, error) {

	var query string
	exercises := []models.Exercises{}
	var err error

	if muscle != "" {
		query = `select * from exercises where (created_by is null or created_by=$1) and muscle_group=$2`
		err = r.db.Select(&exercises, query, id, muscle)

	} else {

		query = `select * from exercises where created_by is null or created_by=$1`
		err = r.db.Select(&exercises, query, id)

	}

	if err != nil {

		return nil, err
	}

	return exercises, nil

}

func (r *ExerciseRepo) InsertExercise(exercise *models.Exercises) error {
	query := `insert into exercises(name,muscle_group,created_by) values ($1,$2,$3)`
	_, err := r.db.Exec(query, exercise.Name, exercise.MuscleGroup, exercise.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}
