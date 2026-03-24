package repository

import (
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
