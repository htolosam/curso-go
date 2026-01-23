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
	repo   *repositories.UserRepository
	config *config.Config
}

func NewUserService(repo *repositories.UserRepository, cfg *config.Config) *UserService {
	return &UserService{repo: repo, config: cfg}
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

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("Credenciales incorrectas: %s", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("Credenciales incorrectas: %s", err)
	}
	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("Error al generar el token: %s", err)
	}
	return token, nil
}

func (s *UserService) generateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}
	// creamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// firmamos el token
	return token.SignedString([]byte(s.config.JwtSecret))
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
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
