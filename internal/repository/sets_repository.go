package repository

import (
	"errors"
	"fmt"
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

func (r *SetsRepo) GetSets(workoutID int) ([]models.Sets, error) {
	var sets []models.Sets
	query := `select * from sets where workout_id=$1 `
	err := r.db.Select(&sets, query, workoutID)
	if err != nil {
		return nil, err
	}
	return sets, err
}

func (r *SetsRepo) DeleteSets(workoutID, setsID int) error {
	fmt.Printf("setsId :%d and workoutI :%d", setsID, workoutID)
	query := `delete from sets  where id=$1 and workout_id=$2`
	result, err := r.db.Exec(query, setsID, workoutID)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("cant find the sets")
	}

	return nil
}
