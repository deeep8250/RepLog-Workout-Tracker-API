package repository

import (
	"database/sql"
	"errors"
	"replog/internal/models"

	"github.com/jmoiron/sqlx"
)

type ExerciseRepo struct {
	db *sqlx.DB
}

func NewExerCiseRepo(Db *sqlx.DB) *ExerciseRepo {
	return &ExerciseRepo{db: Db}
}

// func (r *ExerciseRepo) GetAllExercises(email string) (*models.Exercises, error) {
// 	exercises := models.Exercises{}
// 	query := `select * from exercises where email=$1`
// 	err := r.db.Get(exercises, query, email)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, errors.New("cant find the user")
// 		}

// 		return nil, err
// 	}

// 	return &exercises, nil

// }

func (r *ExerciseRepo) GetAllExercises(id int) ([]models.Exercises, error) {
	exercises := []models.Exercises{}
	query := `select * from exercises where created_by is null or created_by=$1`
	err := r.db.Select(&exercises, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return exercises, nil

}
