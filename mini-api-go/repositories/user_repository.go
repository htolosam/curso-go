package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mini-api-go/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(cxt context.Context, user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(cxt, query, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Error al insertar usuario: %s", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el id del usuario insertado: %s", err)
	}
	user.ID = uint(id)
	return nil
}

func (r *UserRepository) GetByID(cxt context.Context, id uint) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	user := &models.User{}
	err := r.db.QueryRowContext(cxt, query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("No existe un usuario con id %d", id)
		}
		return nil, fmt.Errorf("Error al obtener el usuario con id %d: %s", id, err)
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(cxt context.Context, email string) (*models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	user := &models.User{}
	err := r.db.QueryRowContext(cxt, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("No existe un usuario con email %s", email)
		}
		return nil, fmt.Errorf("Error al obtener el usuario con email %s: %s", email, err)
	}
	return user, nil
}

func (r *UserRepository) EmailExists(cxt context.Context, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	var count int
	err := r.db.QueryRowContext(cxt, query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Error al verificar si existe un usuario con email %s: %s", email, err)
	}
	return count > 0, nil
}
