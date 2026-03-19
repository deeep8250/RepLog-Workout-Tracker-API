package services

import "replog/internal/models"

type AuthServiceInterface interface {
	Register(req *models.RegisterRequest) error
	Login(req *models.LoginRequest) (string, error)
}
