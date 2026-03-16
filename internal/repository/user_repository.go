package repository

import (
	"replog/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `insert into users (email,password_hash) values ($1,$2)`
	_, err := r.db.Exec(query, user.Email, user.Password_hash)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `select * from users  where email=$1`
	err := r.db.Get(user, query, email)
	return user, err
}
