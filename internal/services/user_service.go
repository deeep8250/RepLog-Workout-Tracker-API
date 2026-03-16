package services

import (
	"errors"
	"replog/internal/models"
	"replog/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(UR *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: UR}
}

func (s *AuthService) Register(req *models.RegisterRequest) error {
	existing, err := s.userRepo.GetUserByEmail(req.Email)
	if existing != nil {
		return errors.New("user already exists")
	}
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:         req.Email,
		Password_hash: string(hash),
	}

	return s.userRepo.CreateUser(&user)

}
