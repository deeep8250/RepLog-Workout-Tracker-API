package services

import (
	"database/sql"
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
	if err != nil && errors.Is(err, sql.ErrNoRows) {
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

func (s *AuthService) Login(req *models.LoginRequest) (string, error) {

	if req.Email == "" || req.Password == "" {
		return "", errors.New("please provide data to all fields")
	}

	exist, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not exist")
		}

		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(exist.Password_hash), []byte(req.Password))
	if err != nil {
		return "", err
	}

	token, err := JwtGenerate(exist.Id)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil

}
