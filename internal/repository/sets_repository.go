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

func (r *SetsRepo) GetLastThreeSets(userID, exerciseID int) ([]models.Sets, error) {

	var sets []models.Sets
	query := `select s.id, s.workout_id, s.exercise_id, s.reps, s.weight_kg from sets as s join workouts w on s.workout_id=w.id where w.user_id=$1 and s.exercise_id=$2 order by w.date desc limit 3`
	err := r.db.Select(&sets, query, userID, exerciseID)
	if err != nil {
		return nil, err
	}

	return sets, nil

}

func (r *SetsRepo) ProgressReport(userID, exerciseID int) ([]models.ProgressEntry, error) {

	var report []models.ProgressEntry
	query := `select  s.reps, s.weight_kg, w.date from sets as s join workouts as w on w.id=s.workout_id where w.user_id=$1 and s.exercise_id=$2`
	err := r.db.Select(&report, query, userID, exerciseID)
	if err != nil {
		return nil, err
	}

	return report, nil

}

func (r *SetsRepo) ShouldIncreaseRepsRepo(useID int) ([]models.Exercises, error) {
	query := `select * from exercises where created_by=$1`

	var values []models.Exercises
	err := r.db.Select(&values, query, useID)
	if err != nil {
		return nil, err
	}
	return values, nil
}
