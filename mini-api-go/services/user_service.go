package services

import (
	"context"
	"fmt"
	"mini-api-go/models"
	"mini-api-go/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SingUp(ctx context.Context, name, lastName, email, password string) (*models.User, error) {
	exists, err := s.repo.EmailExists(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("Ya existe un usuario con email %s", email)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error al hashear la contraseña: %s", err)
	}

	user := &models.User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: string(passwordHash),
	}
	if err = s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
