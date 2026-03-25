package repository

import (
	"errors"
	"replog/internal/models"

	"github.com/jmoiron/sqlx"
)

type SetsRepo struct {
	db *sqlx.DB
}

func NewSetsRepo(Db *sqlx.DB) *SetsRepo {
	return &SetsRepo{
		db: Db,
	}
}

func (r *SetsRepo) CreateSets(sets *models.Sets) error {
	query := `insert into sets(workout_id,exercise_id,reps,duration,weight_kg) values ($1,$2,$3,$4,$5)`
	result, err := r.db.Exec(query, sets.WorkoutID, sets.ExerciseID, sets.Reps, sets.Duration, sets.Weight_kg)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cant crete the sets something went wrong ! ")
	}

	return nil

}
