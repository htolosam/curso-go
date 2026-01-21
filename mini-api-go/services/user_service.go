package services

import (
	"context"
	"fmt"
	"mini-api-go/config"
	"mini-api-go/models"
	"mini-api-go/repositories"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SingUp(ctx context.Context, name, lastName, email, password string) (*models.User, error) {
	if err := ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := ValidatePassword(password); err != nil {
		return nil, err
	}
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

func (s *UserService) generateToken(ctx context.Context, userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}
	// creamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// firmamos el token
	return token.SignedString([]byte(config.AppConfig.JwtSecret))
}

func ValidateEmail(email string) error {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegexp.MatchString(email) {
		return fmt.Errorf("email invalido")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("La contraseña debe tener al menos 6 caracteres")
	}
	return nil
}
