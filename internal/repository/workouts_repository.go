package repository

import (
	"errors"
	"replog/internal/models"

	"github.com/jmoiron/sqlx"
)

type WorkoutsRepo struct {
	db *sqlx.DB
}

func NewWorkoutsRepo(Db *sqlx.DB) *WorkoutsRepo {
	return &WorkoutsRepo{db: Db}
}

func (r *WorkoutsRepo) CreateWorkouts(workouts *models.Workouts) error {
	query := `insert into workouts(user_id,notes) values($1,$2) `
	_, err := r.db.Exec(query, workouts.UserID, workouts.Notes)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkoutsRepo) GetWorkouts(userID int) ([]models.Workouts, error) {
	workouts := []models.Workouts{}
	query := `select *  from workouts where user_id=$1`
	err := r.db.Select(&workouts, query, userID)
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

func (r *WorkoutsRepo) GetWorkoutsByID(ID, userID int) (*models.Workouts, error) {
	var workout models.Workouts
	query := `select *  from workouts where id=$1 and user_id=$2`
	err := r.db.Get(&workout, query, ID, userID)
	if err != nil {
		return nil, err
	}
	return &workout, nil
}

func (r *WorkoutsRepo) DeleteWorkoutByID(ID, userID int) error {
	query := `delete from workouts where id=$1 and user_id=$2`
	result, err := r.db.Exec(query, ID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("workout not found")
	}
	return nil
}
